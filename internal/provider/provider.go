package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/prodxcloud/terraform-provider-prodxcloud/internal/resources/s3"
	"github.com/prodxcloud/terraform-provider-prodxcloud/internal/resources/ec2"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"region": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("AWS_REGION", nil),
				Description: "The AWS region to use",
			},
			"access_key": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("AWS_ACCESS_KEY_ID", nil),
				Description: "The AWS access key",
			},
			"secret_key": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("AWS_SECRET_ACCESS_KEY", nil),
				Description: "The AWS secret key",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"prodxcloud_s3_bucket": s3.ResourceBucket(),
			"prodxcloud_ec2_instance": ec2.ResourceInstance(),
		},
		ConfigureContextFunc: providerConfigure,
	}
} 