package com.wgtwo.api.rest.provision;

import com.wgtwo.api.rest.provision.model.ProvisionResponse;
import kong.unirest.HttpResponse;
import kong.unirest.Unirest;

public class ServiceClient {
    private final String baseUri;
    private final String bssId;
    private final String clientId;
    private final String clientSecret;

    ServiceClient(String baseUri, String bssId, String clientId, String clientSecret) {
        this.baseUri = baseUri;
        this.bssId = bssId;
        this.clientId = clientId;
        this.clientSecret = clientSecret;
    }

    public enum Action { ADD, REMOVE }

    public enum Service {
        DATA,
        DATA_HIGHSPEED,
        ROAMING,
        ROAMING_DATA,
    }

    public Data data() { return new Data(); }
    public Roaming roaming() { return new Roaming(); }

    public ProvisionResponse update(String msisdn, String userId, Action action, Service service, String config) {
        UpdateRequest body = new UpdateRequest(bssId, msisdn, userId, action, service, config);
        return Unirest.post(baseUri + "/provision/v1/update")
                .basicAuth(clientId, clientSecret)
                .header("Content-Type", "application/json")
                .body(body)
                .asObject(ProvisionResponse.class)
                .ifFailure(this::throwException)
                .getBody();
    }

    public class Data {
        public ProvisionResponse enable(String msisdn, String userId) {
            return update(msisdn, userId, Action.ADD, Service.DATA, null);
        }

        public ProvisionResponse disable(String msisdn, String userId) {
            return update(msisdn, userId, Action.REMOVE, Service.DATA, null);
        }
    }

    public class Roaming {
        public ProvisionResponse enable(String msisdn, String userId) {
            return update(msisdn, userId, Action.ADD, Service.ROAMING_DATA, null);
        }

        public ProvisionResponse disable(String msisdn, String userId) {
            return update(msisdn, userId, Action.REMOVE, Service.ROAMING_DATA, null);
        }
    }

    private void throwException(HttpResponse<ProvisionResponse> response) {
        if (response.getParsingError().isPresent()) {
            throw new RuntimeException("status=" + response.getStatus() + " content=" + response.getParsingError().get().getOriginalBody());
        }
        throw new RuntimeException("status: " + response.getStatus() + " request ID=" + response.getBody().getRequestId());
    }
}
