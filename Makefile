update_user_kitex_gen:
	kitex -module github.com/gwen0x4c3/easy_note -I idl/ -type protobuf idl/user.proto

update_note_kitex_gen:
	kitex -module github.com/gwen0x4c3/easy_note idl/note.thrift

gen_user_rpc:
	mkdir -p rpc/user && cd rpc/user && \
	kitex -module github.com/gwen0x4c3/easy_note -service userservice \
	 -use github.com/gwen0x4c3/easy_note/kitex_gen \
	 -type protobuf ../../idl/user.proto

gen_note_rpc:
	mkdir -p rpc/note && cd rpc/note && \
	kitex -module github.com/gwen0x4c3/easy_note -service noteservice \
	 -use github.com/gwen0x4c3/easy_note/kitex_gen \
	 ../../idl/note.thrift