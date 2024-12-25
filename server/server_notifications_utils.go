package main

import (
	pb_man "simpleGRPC/proto_defs/manager"
)

func (s *Server) Subscribe(req *pb_man.Notification, stream pb_man.ManagerAsset_SubscribeServer) error {

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
