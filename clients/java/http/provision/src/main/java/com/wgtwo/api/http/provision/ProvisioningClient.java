package com.wgtwo.api.http.provision;

import com.wgtwo.api.common.Environment;
import com.wgtwo.api.http.provision.UpdateRequest.Action;
import com.wgtwo.api.http.provision.UpdateRequest.Service;
import kong.unirest.HttpResponse;
import kong.unirest.Unirest;

public class ProvisioningClient {
    private final String baseUri;
    private final String bssId;
    private final String clientId;
    private final String clientSecret;

    public ProvisioningClient(String baseUri, String bssId, String clientId, String clientSecret) {
        this.baseUri = baseUri;
        this.bssId = bssId;
        this.clientId = clientId;
        this.clientSecret = clientSecret;
    }

    public ProvisioningClient(Environment environment, String bssId, String clientId, String clientSecret) {
        this("https://" + environment.host, bssId, clientId, clientSecret);
    }

    public ProvisionResponse enableRoamingData(String msisdn, String userId) {
        return update(msisdn, userId, Service.ROAMING_DATA, Action.ADD, null);
    }

    public ProvisionResponse disableRoamingData(String msisdn, String userId) {
        return update(msisdn, userId, Service.ROAMING_DATA, Action.REMOVE, null);
    }

    private ProvisionResponse update(String msisdn, String userId, Service service, Action action, String config) {
        UpdateRequest body = new UpdateRequest(bssId, msisdn, userId, action, service, config);
        return Unirest.post(baseUri + "/provision/v1/update")
                .basicAuth(clientId, clientSecret)
                .header("Content-Type", "application/json")
                .body(body)
                .asObject(ProvisionResponse.class)
                .ifFailure(this::throwException)
                .getBody();
    }

    private void throwException(HttpResponse<ProvisionResponse> response) {
        if (response.getParsingError().isPresent()) {
            throw new RuntimeException("status=" + response.getStatus() + " content=" + response.getParsingError().get().getOriginalBody());
        }
        throw new RuntimeException("status: " + response.getStatus() + " request ID=" + response.getBody().getRequestId());
    }
}
