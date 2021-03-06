// Copyright 2017 The Kubernetes Dashboard Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package persistentvolume

import (
	"reflect"
	"testing"

	"github.com/kubernetes/dashboard/src/app/backend/api"
	"github.com/kubernetes/dashboard/src/app/backend/resource/dataselect"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/pkg/api/v1"
)

func TestGetPersistentVolumeList(t *testing.T) {
	cases := []struct {
		persistentVolumes []v1.PersistentVolume
		expected          *PersistentVolumeList
	}{
		{nil, &PersistentVolumeList{Items: []PersistentVolume{}}},
		{
			[]v1.PersistentVolume{
				{
					ObjectMeta: metaV1.ObjectMeta{Name: "foo"},
					Spec: v1.PersistentVolumeSpec{
						PersistentVolumeReclaimPolicy: v1.PersistentVolumeReclaimRecycle,
						AccessModes:                   []v1.PersistentVolumeAccessMode{v1.ReadWriteOnce},
						ClaimRef: &v1.ObjectReference{
							Name:      "myclaim-name",
							Namespace: "default",
						},
						Capacity: nil,
					},
					Status: v1.PersistentVolumeStatus{
						Phase:  v1.VolumePending,
						Reason: "my-reason",
					},
				},
			},
			&PersistentVolumeList{
				ListMeta: api.ListMeta{TotalItems: 1},
				Items: []PersistentVolume{{
					TypeMeta:    api.TypeMeta{Kind: "persistentvolume"},
					ObjectMeta:  api.ObjectMeta{Name: "foo"},
					Capacity:    nil,
					AccessModes: []v1.PersistentVolumeAccessMode{v1.ReadWriteOnce},
					Status:      v1.VolumePending,
					Claim:       "default/myclaim-name",
					Reason:      "my-reason",
				}},
			},
		},
	}
	for _, c := range cases {
		actual := getPersistentVolumeList(c.persistentVolumes, dataselect.NoDataSelect)
		if !reflect.DeepEqual(actual, c.expected) {
			t.Errorf("getPersistentVolumeList(%#v) == \n%#v\nexpected \n%#v\n",
				c.persistentVolumes, actual, c.expected)
		}
	}
}
