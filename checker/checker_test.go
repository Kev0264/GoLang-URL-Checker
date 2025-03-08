package checker

import (
	"testing"
)

func TestCheckURL(t *testing.T) {
	tests:=[]struct{
	name string 
	url string 
	wantStatus string 
	wantErr bool
	}{
		{"Valid URL","https://www.google.com","200 OK",false},
		{"Invalid URL","https://www.google.com/invalid","404 Not Found",false},
		{"Invalid Scheme","ftp://www.google.com","",true},
	}
	for _,tt:=range tests{
		t.Run(tt.name,func(t *testing.T){
			res:=CheckURL(tt.url)
			if tt.wantErr{
				if res.Error==nil{
					t.Errorf("CheckURL(%s) got nil error, want error",tt.url)
				}
				return
			}
			if res.Status!=tt.wantStatus{
				t.Errorf("CheckURL(%s) got status %q, want %q",tt.url,res.Status,tt.wantStatus)
			}
		})
	}
}