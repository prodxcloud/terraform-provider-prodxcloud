package s3

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func ResourceBucket() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceBucketCreate,
		ReadContext:   resourceBucketRead,
		UpdateContext: resourceBucketUpdate,
		DeleteContext: resourceBucketDelete,

		Schema: map[string]*schema.Schema{
			"bucket": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"acl": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "private",
			},
			"versioning": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func resourceBucketCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*s3.Client)
	bucket := d.Get("bucket").(string)

	input := &s3.CreateBucketInput{
		Bucket: &bucket,
		ACL:    s3.BucketCannedACL(d.Get("acl").(string)),
	}

	_, err := client.CreateBucket(ctx, input)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating S3 bucket: %s", err))
	}

	d.SetId(bucket)

	if d.Get("versioning").(bool) {
		_, err := client.PutBucketVersioning(ctx, &s3.PutBucketVersioningInput{
			Bucket: &bucket,
			VersioningConfiguration: &s3.VersioningConfiguration{
				Status: s3.BucketVersioningStatusEnabled,
			},
		})
		if err != nil {
			return diag.FromErr(fmt.Errorf("error enabling versioning: %s", err))
		}
	}

	return resourceBucketRead(ctx, d, meta)
}

func resourceBucketRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*s3.Client)
	bucket := d.Id()

	_, err := client.HeadBucket(ctx, &s3.HeadBucketInput{
		Bucket: &bucket,
	})
	if err != nil {
		return diag.FromErr(fmt.Errorf("error reading S3 bucket: %s", err))
	}

	return nil
}

func resourceBucketUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*s3.Client)
	bucket := d.Id()

	if d.HasChange("versioning") {
		status := s3.BucketVersioningStatusDisabled
		if d.Get("versioning").(bool) {
			status = s3.BucketVersioningStatusEnabled
		}

		_, err := client.PutBucketVersioning(ctx, &s3.PutBucketVersioningInput{
			Bucket: &bucket,
			VersioningConfiguration: &s3.VersioningConfiguration{
				Status: status,
			},
		})
		if err != nil {
			return diag.FromErr(fmt.Errorf("error updating versioning: %s", err))
		}
	}

	return resourceBucketRead(ctx, d, meta)
}

func resourceBucketDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*s3.Client)
	bucket := d.Id()

	_, err := client.DeleteBucket(ctx, &s3.DeleteBucketInput{
		Bucket: &bucket,
	})
	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting S3 bucket: %s", err))
	}

	return nil
} 