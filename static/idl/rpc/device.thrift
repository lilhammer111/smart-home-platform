namespace go device

struct DeviceData {
    1: optional i32 Id;
    2: required string SerialNo;
    3: required string Name;
    4: required string Model;
    5: required string Type;
    6: required string Location;
    7: required string HardwareVersion;
    8: required string SoftwareVersion;
}

struct DeviceTypeReq {
    1: required string Type;
    2: required list<string> Model;
}

struct DeviceFilter {
    1: optional bool IsOnline;
    2: optional i16 Page;
    3: optional i16 Limit;
    4: optional string Sort;
}

service Device {
    i32 RegisterNewDeviceType(1: DeviceTypeReq req);
    DeviceData FindDeviceByID(1: i32 req);
    list<DeviceData> ListDeviceWithFilter(1: DeviceFilter req);
}

service DeviceStatus {

}

service DeviceAlert {

}
