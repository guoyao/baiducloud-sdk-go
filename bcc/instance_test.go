package bcc

import (
	"testing"

	"github.com/baidu/baiducloud-sdk-go/billing"
)

func TestCreateInstance(t *testing.T) {
	bccClient := NewTestClient()
	args := &CreateInstanceArgs{
		Name:    "sdk-test",
		ImageID: "m-Sr1bsnee",
		Billing: billing.Billing{
			PaymentTiming: "Postpaid",
		},
		CPUCount:           1,
		MemoryCapacityInGB: 1,
		PurchaseCount:      1,
	}
	intances, err := bccClient.CreateInstances(args, nil)
	if err != nil {
		t.Errorf("Failed to create image, err: %+v", err)
	}
	t.Logf("Created instances: %+v", intances)

}

func TestDeleteIntsance(t *testing.T) {
	bccClient := NewTestClient()
	InstanceID := "i-eoWGc90b"
	err := bccClient.DeleteInstance(InstanceID, nil)
	if err != nil {
		t.Errorf("Failed to delete instance: %+v, err: %+v", InstanceID, err)
	} else {
		t.Logf("delete instances %+v successfully", InstanceID)
	}
}
