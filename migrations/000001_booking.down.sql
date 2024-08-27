syntax = "proto3";

package gym;

message BookingPersonal {
  string id = 1;
  string user_id = 2;
  string subscription_id = 3;
  int32 payment = 4;
  string access_status = 5;
  string start_date = 6;
  int32 count = 7;
  string created_at = 8;
  string updated_at = 9;
  int64 deleted_at = 10;
}

message BookingGroup {
  string id = 1;
  string user_id = 2;
  string subscription_id = 3;
  int32 payment = 4;
  string access_status = 5;
  string start_date = 6;
  int32 count = 7;
  string created_at = 8;
  string updated_at = 9;
  int64 deleted_at = 10;
}

message BookingCoach {
  string id = 1;
  string user_id = 2;
  string subscription_id = 3;
  int32 payment = 4;
  string access_status = 5;
  string start_date = 6;
  int32 count = 7;
  string created_at = 8;
  string updated_at = 9;
  int64 deleted_at = 10;
}

message CreateBookingPersonalRequest {
  BookingPersonal booking_personal = 1;
}

message GetBookingPersonalRequest {
  string id = 1;
}

message UpdateBookingPersonalRequest {
  BookingPersonal booking_personal = 1;
}

message DeleteBookingPersonalRequest {
  string id = 1;
}

message ListBookingPersonalRequest {
  string user_id = 1;
}

message ListBookingPersonalResponse {
  repeated BookingPersonal booking_personal = 1;
}

message CreateBookingGroupRequest {
  BookingGroup booking_group = 1;
}

message GetBookingGroupRequest {
  string id = 1;
}

message UpdateBookingGroupRequest {
  BookingGroup booking_group = 1;
}

message DeleteBookingGroupRequest {
  string id = 1;
}

message ListBookingGroupRequest {
  string user_id = 1;
}

message ListBookingGroupResponse {
  repeated BookingGroup booking_group = 1;
}

message CreateBookingCoachRequest {
  BookingCoach booking_coach = 1;
}

message GetBookingCoachRequest {
  string id = 1;
}

message UpdateBookingCoachRequest {
  BookingCoach booking_coach = 1;
}

message DeleteBookingCoachRequest {
  string id = 1;
}

message ListBookingCoachRequest {
  string user_id = 1;
}

message ListBookingCoachResponse {
  repeated BookingCoach booking_coach = 1;
}

service BookingPersonalService {
  rpc CreateBookingPersonal (CreateBookingPersonalRequest) returns (BookingPersonal);
  rpc GetBookingPersonal (GetBookingPersonalRequest) returns (BookingPersonal);
  rpc UpdateBookingPersonal (UpdateBookingPersonalRequest) returns (BookingPersonal);
  rpc DeleteBookingPersonal (DeleteBookingPersonalRequest) returns (google.protobuf.Empty);
  rpc ListBookingPersonal (ListBookingPersonalRequest) returns (ListBookingPersonalResponse);
}

service BookingGroupService {
  rpc CreateBookingGroup (CreateBookingGroupRequest) returns (BookingGroup);
  rpc GetBookingGroup (GetBookingGroupRequest) returns (BookingGroup);
  rpc UpdateBookingGroup (UpdateBookingGroupRequest) returns (BookingGroup);
  rpc DeleteBookingGroup (DeleteBookingGroupRequest) returns (google.protobuf.Empty);
  rpc ListBookingGroup (ListBookingGroupRequest) returns (ListBookingGroupResponse);
}

service BookingCoachService {
  rpc CreateBookingCoach (CreateBookingCoachRequest) returns (BookingCoach);
  rpc GetBookingCoach (GetBookingCoachRequest) returns (BookingCoach);
  rpc UpdateBookingCoach (UpdateBookingCoachRequest) returns (BookingCoach);
  rpc DeleteBookingCoach (DeleteBookingCoachRequest) returns (google.protobuf.Empty);
  rpc ListBookingCoach (ListBookingCoachRequest) returns (ListBookingCoachResponse);
}