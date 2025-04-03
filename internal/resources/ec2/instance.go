package ec2

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func ResourceInstance() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceInstanceCreate,
		ReadContext:   resourceInstanceRead,
		UpdateContext: resourceInstanceUpdate,
		DeleteContext: resourceInstanceDelete,

		Schema: map[string]*schema.Schema{
			"ami": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"instance_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"key_name": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceInstanceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*ec2.Client)

	input := &ec2.RunInstancesInput{
		ImageId:      &d.Get("ami").(string),
		InstanceType: types.InstanceType(d.Get("instance_type").(string)),
		MinCount:     int32(1),
		MaxCount:     int32(1),
		SubnetId:     &d.Get("subnet_id").(string),
	}

	if v, ok := d.GetOk("key_name"); ok {
		input.KeyName = &v.(string)
	}

	result, err := client.RunInstances(ctx, input)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating EC2 instance: %s", err))
	}

	instance := result.Instances[0]
	d.SetId(*instance.InstanceId)

	if tags := d.Get("tags").(map[string]interface{}); len(tags) > 0 {
		tagSpecs := []types.TagSpecification{
			{
				ResourceType: types.ResourceTypeInstance,
				Tags:        convertTags(tags),
			},
		}

		_, err := client.CreateTags(ctx, &ec2.CreateTagsInput{
			Resources: []string{d.Id()},
			Tags:      convertTags(tags),
		})
		if err != nil {
			return diag.FromErr(fmt.Errorf("error tagging instance: %s", err))
		}
	}

	return resourceInstanceRead(ctx, d, meta)
}

func resourceInstanceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*ec2.Client)

	input := &ec2.DescribeInstancesInput{
		InstanceIds: []string{d.Id()},
	}

	result, err := client.DescribeInstances(ctx, input)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error reading EC2 instance: %s", err))
	}

	if len(result.Reservations) == 0 || len(result.Reservations[0].Instances) == 0 {
		d.SetId("")
		return nil
	}

	instance := result.Reservations[0].Instances[0]

	d.Set("ami", instance.ImageId)
	d.Set("instance_type", instance.InstanceType)
	d.Set("subnet_id", instance.SubnetId)
	d.Set("key_name", instance.KeyName)

	tags := make(map[string]string)
	for _, tag := range instance.Tags {
		tags[*tag.Key] = *tag.Value
	}
	d.Set("tags", tags)

	return nil
}

func resourceInstanceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*ec2.Client)

	if d.HasChange("instance_type") {
		_, err := client.ModifyInstanceAttribute(ctx, &ec2.ModifyInstanceAttributeInput{
			InstanceId:     &d.Id(),
			InstanceType:   &types.AttributeValue{Value: &d.Get("instance_type").(string)},
		})
		if err != nil {
			return diag.FromErr(fmt.Errorf("error updating instance type: %s", err))
		}
	}

	if d.HasChange("tags") {
		oldTags, newTags := d.GetChange("tags")
		oldTagsMap := oldTags.(map[string]interface{})
		newTagsMap := newTags.(map[string]interface{})

		// Remove old tags
		if len(oldTagsMap) > 0 {
			_, err := client.DeleteTags(ctx, &ec2.DeleteTagsInput{
				Resources: []string{d.Id()},
				Tags:      convertTags(oldTagsMap),
			})
			if err != nil {
				return diag.FromErr(fmt.Errorf("error removing tags: %s", err))
			}
		}

		// Add new tags
		if len(newTagsMap) > 0 {
			_, err := client.CreateTags(ctx, &ec2.CreateTagsInput{
				Resources: []string{d.Id()},
				Tags:      convertTags(newTagsMap),
			})
			if err != nil {
				return diag.FromErr(fmt.Errorf("error adding tags: %s", err))
			}
		}
	}

	return resourceInstanceRead(ctx, d, meta)
}

func resourceInstanceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*ec2.Client)

	_, err := client.TerminateInstances(ctx, &ec2.TerminateInstancesInput{
		InstanceIds: []string{d.Id()},
	})
	if err != nil {
		return diag.FromErr(fmt.Errorf("error terminating instance: %s", err))
	}

	return nil
}

func convertTags(tags map[string]interface{}) []types.Tag {
	result := make([]types.Tag, 0, len(tags))
	for k, v := range tags {
		result = append(result, types.Tag{
			Key:   &k,
			Value: &v.(string),
		})
	}
	return result
} 