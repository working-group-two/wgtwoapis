package com.wgtwo.api.rest.provision;

class UpdateRequest {
    String bssid;
    String msisdn;
    String userid;
    ServiceBody service = new ServiceBody();

    UpdateRequest(String bssid, String msisdn, String userid, Action action, Service serviceName, String config) {
        this.bssid = bssid;
        this.msisdn = msisdn;
        this.userid = userid;
        this.service.action = action;
        this.service.name = serviceName;
        this.service.config = config;
    }

    static class ServiceBody {
        Action action;
        Service name;
        String config;
    }

    enum Action {ADD, REMOVE}

    enum Service {ROAMING_DATA}
}
