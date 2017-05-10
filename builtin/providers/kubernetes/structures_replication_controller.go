package kubernetes

import (
	"k8s.io/kubernetes/pkg/api/v1"
)

func flattenReplicationControllerSpec(in v1.ReplicationControllerSpec) []interface{} {
	att := make(map[string]interface{})
	if in.MinReadySeconds != 0 {
		att["min_ready_seconds"] = in.MinReadySeconds
	}
	if in.Replicas != nil {
		att["replicas"] = *in.Replicas
	}
	if len(in.Selector) > 0 {
		att["selector"] = in.Selector
	}
	att["template"] = flattenPodSpec(in.Template.Spec)

	return []interface{}{att}
}

func expandReplicationControllerSpec(rc []interface{}) v1.ReplicationControllerSpec {
	obj := v1.ReplicationControllerSpec{}
	if len(p) == 0 || p[0] == nil {
		return obj, nil
	}
	in := rc[0].(map[string]interface{})

	if v, ok := in["min_ready_seconds"].(int); ok {
		obj.MinReadySeconds = int32(v)
	}
	if v, ok := in["replicas"].(int); ok {
		obj.Replicas = ptrToInt32(v)
	}
	if v, ok := in["selector"].(map[string]interface{}); ok {
		obj.Selector = v
	}
	if v, ok := in["template"].([]interface{}); ok {
		obj.Template = v1.PodTemplateSpec{
			Spec: expandPodSpec(v),
		}
	}

	return obj
}
