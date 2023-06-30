// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap/confmaptest"
)

func TestResourceAttributesConfig(t *testing.T) {
	tests := []struct {
		name string
		want ResourceAttributesConfig
	}{
		{
			name: "default",
			want: DefaultResourceAttributesConfig(),
		},
		{
			name: "all_set",
			want: ResourceAttributesConfig{
				CloudAccountID:        ResourceAttributeConfig{Enabled: true},
				CloudAvailabilityZone: ResourceAttributeConfig{Enabled: true},
				CloudPlatform:         ResourceAttributeConfig{Enabled: true},
				CloudProvider:         ResourceAttributeConfig{Enabled: true},
				CloudRegion:           ResourceAttributeConfig{Enabled: true},
				HostID:                ResourceAttributeConfig{Enabled: true},
				HostImageID:           ResourceAttributeConfig{Enabled: true},
				HostName:              ResourceAttributeConfig{Enabled: true},
				HostType:              ResourceAttributeConfig{Enabled: true},
			},
		},
		{
			name: "none_set",
			want: ResourceAttributesConfig{
				CloudAccountID:        ResourceAttributeConfig{Enabled: false},
				CloudAvailabilityZone: ResourceAttributeConfig{Enabled: false},
				CloudPlatform:         ResourceAttributeConfig{Enabled: false},
				CloudProvider:         ResourceAttributeConfig{Enabled: false},
				CloudRegion:           ResourceAttributeConfig{Enabled: false},
				HostID:                ResourceAttributeConfig{Enabled: false},
				HostImageID:           ResourceAttributeConfig{Enabled: false},
				HostName:              ResourceAttributeConfig{Enabled: false},
				HostType:              ResourceAttributeConfig{Enabled: false},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
			require.NoError(t, err)
			sub, err := cm.Sub(tt.name)
			require.NoError(t, err)
			sub, err = sub.Sub("resource_attributes")
			require.NoError(t, err)
			cfg := DefaultResourceAttributesConfig()
			require.NoError(t, component.UnmarshalConfig(sub, &cfg))

			if diff := cmp.Diff(tt.want, cfg, cmpopts.IgnoreUnexported(ResourceAttributeConfig{})); diff != "" {
				t.Errorf("Config mismatch (-expected +actual):\n%s", diff)
			}
		})
	}
}