package com.wgtwo.api.util.phonenumber;

import com.google.i18n.phonenumbers.NumberParseException;
import com.google.i18n.phonenumbers.PhoneNumberUtil;
import com.google.i18n.phonenumbers.PhoneNumberUtil.PhoneNumberFormat;
import com.google.i18n.phonenumbers.Phonenumber.PhoneNumber;
import com.wgtwo.api.common.v0.PhoneNumberProto;

public final class PhoneNumbers {
    private static final PhoneNumberUtil util = PhoneNumberUtil.getInstance();

    private PhoneNumbers() {}

    public static PhoneNumber of(String number, String defaultRegion) {
        return parse(number, defaultRegion);
    }

    public static PhoneNumber of(String number) {
        return of(number, "ZZ");
    }

    public static PhoneNumber of(PhoneNumberProto.PhoneNumber phoneNumberProto) {
        return of(phoneNumberProto.getE164());
    }

    public static String asE164(PhoneNumber phoneNumber) {
        return util.format(phoneNumber, PhoneNumberFormat.E164);
    }

    public static String asInternational(PhoneNumber phoneNumber) {
        return util.format(phoneNumber, PhoneNumberFormat.INTERNATIONAL);
    }

    public static PhoneNumberProto.PhoneNumber toPhoneNumberProto(PhoneNumber phoneNumber) {
        return PhoneNumberProto.PhoneNumber.newBuilder()
                .setE164(asE164(phoneNumber))
                .build();
    }

    public static String getRegionCode(PhoneNumber phoneNumber) {
        return util.getRegionCodeForNumber(phoneNumber);
    }

    public static Boolean isValid(PhoneNumber phoneNumber) {
        return util.isValidNumber(phoneNumber);
    }

    private static PhoneNumber parse(String number, String defaultRegion) {
        try {
            return util.parse(number, defaultRegion);
        } catch (NumberParseException e) {
            return null;
        }
    }
}
