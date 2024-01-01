HADOOP_COMMON_PROTOS = $(shell find internal/proto/hadoop_common -name '*.proto')
HADOOP_HDFS_PROTOS = $(shell find internal/proto/hadoop_hdfs -name '*.proto')
GENERATED_PROTOS = $(shell echo "$(HADOOP_HDFS_PROTOS) $(HADOOP_COMMON_PROTOS)" | sed 's/\.proto/\.pb\.go/g')
SOURCES = $(shell find . -name '*.go') $(GENERATED_PROTOS)

# Protobuf needs one of these for every 'import "foo.proto"' in .protoc files.
PROTO_MAPPING = MSecurity.proto=github.com/dmpdl/ghdfs

RELEASE_NAME = ghdfs-$(TAG)-$(ARCH)

proto:
	protoc --go_out=internal/proto/hadoop_common \
		--go_opt=paths=source_relative \
		-Iinternal/proto/hadoop_common \
		-Iinternal/proto/hadoop_hdfs \
		$(HADOOP_COMMON_PROTOS)
	protoc --go_out=internal/proto/hadoop_hdfs \
		--go_opt=paths=source_relative \
		-Iinternal/proto/hadoop_common \
		-Iinternal/proto/hadoop_hdfs \
		$(HADOOP_HDFS_PROTOS)

