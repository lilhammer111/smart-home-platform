namespace go pet_feeder

struct DeviceData {
    1: optional i32 Id;
    1: required string SerialNo;
    2: required string Name;
    3: required string Model;
    4: required string Type;
    5: required string Location;
    6: required string HardwareVersion;
    6: required string SoftwareVersion;
}


service DeviceSrv {
    DeviceData FindDeviceByID(1: i32 req);
    list<DeviceData> FilterDevices (1:  )
}

service StatusSrv {

}

service AlertSrv {

}

service FeedingSrv {

}