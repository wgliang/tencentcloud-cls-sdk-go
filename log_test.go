package cls

import (
	"testing"
	"time"
)

func TestClSCleint_UploadLog(t *testing.T) {
	k1, v1, k2, v2 := "key1", "value1", "key2", "value2"
	t1 := time.Now().Unix()

	type fields struct {
		SecretId  string
		SecretKey string
		Host      string
	}
	type args struct {
		logTopicID   string
		logGroupList LogGroupList
		hash         string
		compress     bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
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
				logGroupList: LogGroupList{
					LogGroupList: []*LogGroup{
						&LogGroup{
							Logs: []*Log{
								&Log{
									Time: &t1,
									Contents: []*Log_Content{
										&Log_Content{
											Key:   &k1,
											Value: &v1,
										},
										&Log_Content{
											Key:   &k2,
											Value: &v2,
										},
									},
								},
							},
						},
					},
				},
				hash:     "",
				compress: false,
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
				logGroupList: LogGroupList{
					LogGroupList: []*LogGroup{
						&LogGroup{
							Logs: []*Log{
								&Log{
									Time: &t1,
									Contents: []*Log_Content{
										&Log_Content{
											Key:   &k1,
											Value: &v1,
										},
										&Log_Content{
											Key:   &k2,
											Value: &v2,
										},
									},
								},
							},
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
			if err := cls.UploadLog(tt.args.logTopicID, tt.args.logGroupList, tt.args.hash, tt.args.compress); (err != nil) != tt.wantErr {
				t.Errorf("ClSCleint.UploadLog() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClSCleint_GetLogStart(t *testing.T) {
	type fields struct {
		SecretId  string
		SecretKey string
		Host      string
	}
	type args struct {
		logTopicID string
		start      string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantCursor string
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
			gotCursor, err := cls.GetLogStart(tt.args.logTopicID, tt.args.start)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClSCleint.GetLogStart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotCursor != tt.wantCursor {
				t.Errorf("ClSCleint.GetLogStart() = %v, want %v", gotCursor, tt.wantCursor)
			}
		})
	}
}

func TestClSCleint_SearchLog(t *testing.T) {
	type fields struct {
		SecretId  string
		SecretKey string
		Host      string
	}
	type args struct {
		requestDataMap map[string]string
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
			got, err := cls.SearchLog(tt.args.requestDataMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClSCleint.SearchLog() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ClSCleint.SearchLog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClSCleint_DowloadLog(t *testing.T) {
	type fields struct {
		SecretId  string
		SecretKey string
		Host      string
	}
	type args struct {
		logTopicID string
		cursor     string
		count      string
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
			if err := cls.DowloadLog(tt.args.logTopicID, tt.args.cursor, tt.args.count); (err != nil) != tt.wantErr {
				t.Errorf("ClSCleint.DowloadLog() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
