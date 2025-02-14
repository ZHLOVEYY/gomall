#位于app/frontend目录下执行
.PHONY:
gen-frontend:
	cwgo server --type HTTP --idl ../../idl/frontend/home.proto --service frontend --module github.com/ZHLOVEYY/gomall/app/frontend -I ../../idl