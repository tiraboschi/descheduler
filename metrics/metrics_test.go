/*
Copyright 2021 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package metrics

import (
	"testing"
)

func TestNormalizePluginName(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"LowNodeUtilization", "low_node_utilization"},
		{"RemovePodsViolatingNodeTaints", "remove_pods_violating_node_taints"},
		{"HighNodeUtilization", "high_node_utilization"},
		{"RemoveDuplicates", "remove_duplicates"},
		{"simple", "simple"},
		{"SimplePlugin", "simple_plugin"},
		{"ABC", "a_b_c"},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := normalizePluginName(test.input)
			if result != test.expected {
				t.Errorf("normalizePluginName(%q) = %q, expected %q", test.input, result, test.expected)
			}
		})
	}
}

func TestGetPluginSubsystem(t *testing.T) {
	tests := []struct {
		pluginName string
		expected   string
	}{
		{"LowNodeUtilization", "descheduler_low_node_utilization"},
		{"RemovePodsViolatingNodeTaints", "descheduler_remove_pods_violating_node_taints"},
		{"DefaultEvictor", "descheduler_default_evictor"},
	}

	for _, test := range tests {
		t.Run(test.pluginName, func(t *testing.T) {
			result := getPluginSubsystem(test.pluginName)
			if result != test.expected {
				t.Errorf("getPluginSubsystem(%q) = %q, expected %q", test.pluginName, result, test.expected)
			}
		})
	}
}

func TestPluginMetricsRegistry_GetPluginSubsystem(t *testing.T) {
	registry := NewPluginMetricsRegistry()

	result := registry.GetPluginSubsystem("LowNodeUtilization")
	expected := "descheduler_low_node_utilization"

	if result != expected {
		t.Errorf("GetPluginSubsystem(%q) = %q, expected %q", "LowNodeUtilization", result, expected)
	}
}
