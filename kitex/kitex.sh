#!/usr/bin/env bash
kitex -module github.com/houqingying/douyin-lite -I ./idl -v -service usersrv user.proto
kitex -module github.com/houqingying/douyin-lite -I ./idl -v -service commentsrv comment.proto
kitex -module github.com/houqingying/douyin-lite -I ./idl -v -service relationsrv relation.proto
kitex -module github.com/houqingying/douyin-lite -I ./idl -v -service favoritesrv favorite.proto
kitex -module github.com/houqingying/douyin-lite -I ./idl -v -service messagesrv message.proto
kitex -module github.com/houqingying/douyin-lite -I ./idl -v -service videosrv video.proto