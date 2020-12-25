package cls

import (
	"reflect"
	"testing"
)

func TestClSCleint_GetLogIndex(t *testing.T) {
	type fields struct {
		SecretId  string
		SecretKey string
		Host      string
	}
	type args struct {
		logTopicID string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantIndex Index
		wantErr   bool
	}{
		{
			name: "test01",
			fields: fields{
				SecretId:  "secret-id01",
				SecretKey: "secret-key01",
				Host:      "127.0.0.1:8080",
			},
			args: args{
				logTopicID: "topic01",
			},
			wantIndex: Index{
				TopicID:   "topic01",
				Effective: false,
				Rule: Rule{
					FullText: FullText{
						CaseSensitive: false,
					},
					KeyValue: KeyValue{
						CaseSensitive: false,
						Keys:          []string{"key1", "key2"},
						Types:         []string{"type1", "type2"},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "test02",
			fields: fields{
				SecretId:  "secret-id01",
				SecretKey: "secret-key01",
				Host:      "127.0.0.1:8080",
			},
			args: args{
				logTopicID: "topic02",
			},
			wantIndex: Index{},
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cls := &ClSCleint{
				SecretId:  tt.fields.SecretId,
				SecretKey: tt.fields.SecretKey,
				Host:      tt.fields.Host,
			}
			gotIndex, err := cls.GetLogIndex(tt.args.logTopicID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClSCleint.GetLogIndex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotIndex, tt.wantIndex) {
				t.Errorf("ClSCleint.GetLogIndex() = %v, want %v", gotIndex, tt.wantIndex)
			}
		})
	}
}

func TestClSCleint_UpdateLogIndex(t *testing.T) {
	type fields struct {
		SecretId  string
		SecretKey string
		Host      string
	}
	type args struct {
		logIndex Index
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test03",
			fields: fields{
				SecretId:  "secret-id01",
				SecretKey: "secret-key01",
				Host:      "127.0.0.1:8080",
			},
			args: args{
				logIndex: Index{
					TopicID:   "topic03",
					Effective: false,
					Rule: Rule{
						FullText: FullText{
							CaseSensitive: false,
						},
						KeyValue: KeyValue{
							CaseSensitive: false,
							Keys:          []string{"key1", "key2"},
							Types:         []string{"type1", "type2"},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "test04",
			fields: fields{
				SecretId:  "secret-id01",
				SecretKey: "secret-key01",
				Host:      "127.0.0.1:8080",
			},
			args: args{
				logIndex: Index{
					TopicID:   "topic02",
					Effective: false,
					Rule: Rule{
						FullText: FullText{
							CaseSensitive: false,
						},
						KeyValue: KeyValue{
							CaseSensitive: false,
							Keys:          []string{"key1", "key2"},
							Types:         []string{"type1", "type2"},
						},
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cls := &ClSCleint{
				SecretId:  tt.fields.SecretId,
				SecretKey: tt.fields.SecretKey,
				Host:      tt.fields.Host,
			}

			if err := cls.UpdateLogIndex(tt.args.logIndex); (err != nil) != tt.wantErr {
				t.Errorf("ClSCleint.UpdateLogIndex() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
