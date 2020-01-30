package com.wgtwo.api.rest.provision;

import java.util.Objects;

public class ProvisionResponse {
    private String requestid;

    public String getRequestId() {
        return requestid;
    }

    @Override
    public String toString() {
        return "ProvisionResponse{requestId='" + requestid + "\'}";
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        ProvisionResponse that = (ProvisionResponse) o;
        return Objects.equals(requestid, that.requestid);
    }

    @Override
    public int hashCode() {
        return Objects.hash(requestid);
    }
}
