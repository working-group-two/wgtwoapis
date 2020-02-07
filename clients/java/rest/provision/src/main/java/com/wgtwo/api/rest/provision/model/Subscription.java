package com.wgtwo.api.rest.provision.model;

import java.time.Instant;
import java.util.Map;
import java.util.Set;
import java.util.StringJoiner;

public class Subscription {
  public long id;
  public String msisdn;
  public boolean active;
  public String userId;
  public String bss;
  public Set<SimcardApiV2> simcards;
  public Map<String, Object> services;
  public Instant created;
  public Instant lastModified;

  @Override
  public String toString() {
    return new StringJoiner(", ", Subscription.class.getSimpleName() + "[", "]")
            .add("id=" + id)
            .add("msisdn='" + msisdn + "'")
            .add("active=" + active)
            .add("userId='" + userId + "'")
            .add("bss='" + bss + "'")
            .add("simcards=" + simcards)
            .add("services=" + services)
            .add("created=" + created)
            .add("lastModified=" + lastModified)
            .toString();
  }

  public static class SimcardApiV2 {
    public Set<String> imsis;
    public String iccid;
    public String bss;
    public Instant created;
    public Instant lastModified;

    @Override
    public String toString() {
      return new StringJoiner(", ", SimcardApiV2.class.getSimpleName() + "[", "]")
              .add("imsis=" + imsis)
              .add("iccid='" + iccid + "'")
              .add("bss='" + bss + "'")
              .add("created=" + created)
              .add("lastModified=" + lastModified)
              .toString();
    }
  }
}
