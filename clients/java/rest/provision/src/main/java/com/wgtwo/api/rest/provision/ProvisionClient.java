package com.wgtwo.api.rest.provision;

import com.wgtwo.api.common.Environment;
import com.wgtwo.api.rest.provision.model.Subscription;
import kong.unirest.Unirest;

public class ProvisionClient {
    private final String baseUri;
    private final String bssId;
    private final String clientId;
    private final String clientSecret;

    public ProvisionClient(String baseUri, String bssId, String clientId, String clientSecret) {
        this.baseUri = baseUri;
        this.bssId = bssId;
        this.clientId = clientId;
        this.clientSecret = clientSecret;
        Unirest.config().setObjectMapper(new JsonObjectMapper());
    }

    public ProvisionClient(Environment environment, String bssId, String clientId, String clientSecret) {
        this("https://" + environment.host, bssId, clientId, clientSecret);
    }

    public ServiceClient services() {
        return new ServiceClient(baseUri, bssId, clientId, clientSecret);
    }

    public Subscription getSubscription(String msisdn) {
        return Unirest.get(baseUri + "/subscription/v1/msisdn/" + msisdn)
                .basicAuth(clientId, clientSecret)
                .header("Content-Type", "application/json")
                .asObject(Subscription.class)
                .ifFailure(response -> {
                    if (response.getParsingError().isPresent()) {
                        System.out.println(response.getParsingError().get());
                        throw new RuntimeException("status=" + response.getStatus() + " content=" + response.getParsingError().get().getOriginalBody());
                    }
                    throw new RuntimeException("status: " + response.getStatus() + " sub=" + response.getBody());
                })
                .getBody();
    }
}
