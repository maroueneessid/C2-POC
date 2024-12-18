package main

import (
	pb "simpleGRPC/proto_defs"
)

func (s *Server) Subscribe(req *pb.Notification, stream pb.AssetService_SubscribeServer) error {

	for {
		select {

		case <-stream.Context().Done():
			return stream.Context().Err()

		case notif := <-s.notifs:
			if err := stream.Send(notif); err != nil {
				return err
			}
		}
	}

}
