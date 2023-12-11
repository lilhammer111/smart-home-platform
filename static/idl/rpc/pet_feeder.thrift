namespace go pet_feeder

struct DeviceData {
    1: optional i32 Id;
    1: required string DeviceSerialNo;
    2: required string DeviceName;
    3: required string DeviceStyle;
    4: required string DeviceType;
    5: required string Location;
    6: required string HardwareVersion;
    6: required string SoftwareVersion;
}


service PetFeederSrv {
    UserData FindDevicesByID(1: i32 req);
    bool VerifyCredentials(1: CredentialsReq req);
    UserData CreateOrUpdateUser(1: UserData req);
    UserData DeleteUser(1: i32 req);
}