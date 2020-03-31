package com.wgtwo.api.util.calling;

import java.util.UUID;

public class CallId {

  public static String generateCallId() {
    return UUID.randomUUID().toString();
  }

}
