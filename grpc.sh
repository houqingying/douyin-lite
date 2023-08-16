#protoc --proto_path=. --go_out=. --go-grpc_out=. protobuf/*.proto
set -eux
cd ./protobuf
array=(`ls *.proto`); \
# get prefix of file name
#array=(`echo $array | sed 's/.proto//g'`);
for item in ${array[@]}; do
  a=(`echo $item | sed 's/.proto//g'`);
  protoc --proto_path=. --go_out=./grpc_gen/$a --go-grpc_out=./grpc_gen/$a --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative "$a".proto

#  protoc --proto_path=. --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative protobuf/"$a".proto; \
done