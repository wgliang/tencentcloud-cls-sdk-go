package cls

import (
	"reflect"
	"testing"
)

func TestClSCleint_CreateMachineGroup(t *testing.T) {
	type fields struct {
		SecretId  string
		SecretKey string
		Host      string
	}
	type args struct {
		machineGroup *MachineGroup
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cls := &ClSCleint{
				SecretId:  tt.fields.SecretId,
				SecretKey: tt.fields.SecretKey,
				Host:      tt.fields.Host,
			}
			got, err := cls.CreateMachineGroup(tt.args.machineGroup)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClSCleint.CreateMachineGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ClSCleint.CreateMachineGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClSCleint_GetMachineGroup(t *testing.T) {
	type fields struct {
		SecretId  string
		SecretKey string
		Host      string
	}
	type args struct {
		groupID string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantGroup MachineGroup
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cls := &ClSCleint{
				SecretId:  tt.fields.SecretId,
				SecretKey: tt.fields.SecretKey,
				Host:      tt.fields.Host,
			}
			gotGroup, err := cls.GetMachineGroup(tt.args.groupID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClSCleint.GetMachineGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotGroup, tt.wantGroup) {
				t.Errorf("ClSCleint.GetMachineGroup() = %v, want %v", gotGroup, tt.wantGroup)
			}
		})
	}
}

func TestClSCleint_GetMachines(t *testing.T) {
	type fields struct {
		SecretId  string
		SecretKey string
		Host      string
	}
	type args struct {
		groupID string
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantMachines Machines
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cls := &ClSCleint{
				SecretId:  tt.fields.SecretId,
				SecretKey: tt.fields.SecretKey,
				Host:      tt.fields.Host,
			}
			gotMachines, err := cls.GetMachines(tt.args.groupID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClSCleint.GetMachines() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotMachines, tt.wantMachines) {
				t.Errorf("ClSCleint.GetMachines() = %v, want %v", gotMachines, tt.wantMachines)
			}
		})
	}
}

func TestClSCleint_GetMachineGroups(t *testing.T) {
	type fields struct {
		SecretId  string
		SecretKey string
		Host      string
	}
	tests := []struct {
		name       string
		fields     fields
		wantGroups MachineGroups
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cls := &ClSCleint{
				SecretId:  tt.fields.SecretId,
				SecretKey: tt.fields.SecretKey,
				Host:      tt.fields.Host,
			}
			gotGroups, err := cls.GetMachineGroups()
			if (err != nil) != tt.wantErr {
				t.Errorf("ClSCleint.GetMachineGroups() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotGroups, tt.wantGroups) {
				t.Errorf("ClSCleint.GetMachineGroups() = %v, want %v", gotGroups, tt.wantGroups)
			}
		})
	}
}

func TestClSCleint_UpdateMachineGroup(t *testing.T) {
	type fields struct {
		SecretId  string
		SecretKey string
		Host      string
	}
	type args struct {
		group *MachineGroup
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cls := &ClSCleint{
				SecretId:  tt.fields.SecretId,
				SecretKey: tt.fields.SecretKey,
				Host:      tt.fields.Host,
			}
			if err := cls.UpdateMachineGroup(tt.args.group); (err != nil) != tt.wantErr {
				t.Errorf("ClSCleint.UpdateMachineGroup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClSCleint_DeleteMachineGroup(t *testing.T) {
	type fields struct {
		SecretId  string
		SecretKey string
		Host      string
	}
	type args struct {
		groupID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cls := &ClSCleint{
				SecretId:  tt.fields.SecretId,
				SecretKey: tt.fields.SecretKey,
				Host:      tt.fields.Host,
			}
			if err := cls.DeleteMachineGroup(tt.args.groupID); (err != nil) != tt.wantErr {
				t.Errorf("ClSCleint.DeleteMachineGroup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
