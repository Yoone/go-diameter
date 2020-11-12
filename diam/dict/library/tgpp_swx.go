package library

var TgppSwx = DictInfo{
	Name: "3GPP Swx",
	XML: `<?xml version="1.0" encoding="UTF-8"?>
<diameter>
    <!--
        3GPP TS 29.273
        http://www.qtc.jp/3GPP/Specs/29273-920.pdf
    -->
    <application id="16777265" type="auth" name="TGPP SWX">
        <vendor id="10415" name="TGPP"/>
        <command code="303" short="MA" name="Multimedia-Authentication">
            <request>
                <!-- http://www.qtc.jp/3GPP/Specs/29273-920.pdf Section 8.2.2.1 -->
                <rule avp="Session-Id" required="true" max="1"/>
                <rule avp="Vendor-Specific-Application-Id" required="true" max="1"/>
                <rule avp="Auth-Session-State" required="true" max="1"/>
                <rule avp="Origin-Host" required="true" max="1"/>
                <rule avp="Origin-Realm" required="true" max="1"/>
                <rule avp="Destination-Host" required="false" max="1"/>
                <rule avp="Destination-Realm" required="true" max="1"/>
                <rule avp="User-Name" required="true" max="1"/>
                <rule avp="RAT-Type" required="false" max="1"/>
                <rule avp="ANID" required="false" max="1"/>
                <rule avp="Visited-Network-Identifier" required="false" max="1" />
                <rule avp="Terminal-Information" required="false" max="1"/>
                <rule avp="SIP-Auth-Data-Item" required="false" max="1"/>
                <rule avp="SIP-Number-Auth-Items" required="false" max="1"/>
                <rule avp="Supported-Features" required="false"/>
                <rule avp="AVP" required="false"/>
            </request>
            <answer>
                <!-- http://www.qtc.jp/3GPP/Specs/29273-920.pdf Section 8.2.2.1 -->
                <rule avp="Session-Id" required="true" max="1"/>
                <rule avp="Vendor-Specific-Application-Id" required="true" max="1"/>
                <rule avp="Result-Code" required="false" max="1"/>
                <rule avp="Experimental-Result" required="false" max="1"/>
                <rule avp="Auth-Session-State" required="true" max="1"/>
                <rule avp="Origin-Host" required="true" max="1"/>
                <rule avp="Origin-Realm" required="true" max="1"/>
                <rule avp="User-Name" required="true" max="1"/>
                <rule avp="SIP-Number-Auth-Items" required="false" max="1"/>
                <rule avp="SIP-Auth-Data-Item" required="false"/>
                <rule avp="TGPP-AAA-Server-Name" required="false" max="1"/>
                <rule avp="Supported-Features" required="false"/>
                <rule avp="AVP" required="false"/>
            </answer>
        </command>
        <command code="301" short="SA" name="Server-Assignment">
            <request>
                <!-- http://www.qtc.jp/3GPP/Specs/29273-920.pdf Section 8.2.2.3 -->
                <rule avp="Session-Id" required="true" max="1"/>
                <rule avp="Vendor-Specific-Application-Id" required="true" max="1"/>
                <rule avp="Auth-Session-State" required="true" max="1"/>
                <rule avp="Origin-Host" required="true" max="1"/>
                <rule avp="Origin-Realm" required="true" max="1"/>
                <rule avp="Destination-Host" required="false" max="1"/>
                <rule avp="Destination-Realm" required="true" max="1"/>
                <rule avp="Service-Selection" required="false" max="1"/>
                <rule avp="Context-Identifier" required="false" max="1"/>
                <rule avp="MIP6-Agent-Info" required="false" max="1"/>
                <rule avp="Visited-Network-Identifier" required="false" max="1"/>
                <rule avp="User-Name" required="true" max="1"/>
                <rule avp="Server-Assignment-Type" required="true" max="1"/>
                <rule avp="Supported-Features" required="false"/>
                <rule avp="AVP" required="false"/>
            </request>
            <answer>
                <!-- http://www.qtc.jp/3GPP/Specs/29273-920.pdf Section 8.2.2.3 -->
                <rule avp="Session-Id" required="true" max="1"/>
                <rule avp="Vendor-Specific-Application-Id" required="true" max="1"/>
                <rule avp="Result-Code" required="false" max="1"/>
                <rule avp="Experimental-Result" required="false" max="1"/>
                <rule avp="Auth-Session-State" required="true" max="1"/>
                <rule avp="Origin-Host" required="true" max="1"/>
                <rule avp="Origin-Realm" required="true" max="1"/>
                <rule avp="User-Name" required="true" max="1"/>
                <rule avp="Non-3GPP-User-Data" required="false" max="1"/>
                <rule avp="TGPP-AAA-Server-Name" required="false" max="1"/>
                <rule avp="Supported-Features" required="false"/>
                <rule avp="AVP" required="false"/>
            </answer>
        </command>
        <command code="304" short="RT" name="Registration-Termination">
            <request>
                <!-- http://www.qtc.jp/3GPP/Specs/29273-920.pdf Section 8.2.2.4 -->
                <rule avp="Session-Id" required="true" max="1"/>
                <rule avp="DRMP" required="false" max="1" />
                <rule avp="Vendor-Specific-Application-Id" required="true" max="1"/>
                <rule avp="Auth-Session-State" required="true" max="1"/>
                <rule avp="Origin-Host" required="true" max="1"/>
                <rule avp="Origin-Realm" required="true" max="1"/>
                <rule avp="Destination-Host" required="false" max="1"/>
                <rule avp="Destination-Realm" required="true" max="1"/>
                <rule avp="User-Name" required="true" max="1"/>
                <rule avp="Deregistration-Reason" required="true" max="1"/>
                <rule avp="Supported-Features" required="false"/>
                <rule avp="AVP" required="false"/>
            </request>
            <answer>
                <!-- http://www.qtc.jp/3GPP/Specs/29273-920.pdf Section 8.2.2.4 -->
                <rule avp="Session-Id" required="true" max="1"/>
                <rule avp="DRMP" required="false" max="1" />
                <rule avp="Vendor-Specific-Application-Id" required="true" max="1"/>
                <rule avp="Result-Code" required="false" max="1"/>
                <rule avp="Experimental-Result" required="false" max="1"/>
                <rule avp="Auth-Session-State" required="true" max="1"/>
                <rule avp="Origin-Host" required="true" max="1"/>
                <rule avp="Origin-Realm" required="true" max="1"/>
                <rule avp="Supported-Features" required="false"/>
                <rule avp="AVP" required="false"/>
            </answer>
        </command>

        <avp name="RAT-Type" code="1032" must="M,V" may="P" may-encrypt="Y" vendor-id="10415">
            <!-- http://www.qtc.jp/3GPP/Specs/29273-920.pdf Section 5.2.3.6 -->
            <data type="Enumerated">
                <item code="0" name="WLAN"/>
                <item code="1" name="VIRTUAL"/>
                <item code="1000" name="UTRAN"/>
                <item code="1001" name="GERAN"/>
                <item code="1002" name="GAN"/>
                <item code="1003" name="HSPA_EVOLUTION"/>
                <item code="1004" name="EUTRAN"/>
                <item code="2000" name="CDMA2000_1X"/>
                <item code="2001" name="HRPD"/>
                <item code="2002" name="UMB"/>
                <item code="2003" name="EHRPD"/>
            </data>
        </avp>

        <avp name="ANID" code="1504" must="M,V" may-encrypt="N" vendor-id="10415">
            <!-- http://www.qtc.jp/3GPP/Specs/29273-920.pdf Section 5.2.3.7 -->
            <data type="UTF8String"/>
        </avp>

        <avp name="Visited-Network-Identifier" code="600" must="M,V" may-encrypt="N" vendor-id="10415">
            <!-- http://www.qtc.jp/3GPP/Specs/29273-920.pdf Section 9.2.3.1.2 -->
            <data type="OctetString"/>
        </avp>

        <avp name="Terminal-Information" code="1401" must="M,V" may-encrypt="N" vendor-id="10415">
            <!-- 3GPP TS 29.272 Section 7.3.3 -->
            <data type="Grouped">
                <rule avp="IMEI" required="false" max="1"/>
                <rule avp="TGPP2-MEID" required="false" max="1"/>
                <rule avp="Software-Version" required="false" max="1"/>
                <rule avp="AVP" required="false"/>
            </data>
        </avp>

        <avp name="SIP-Auth-Data-Item" code="612" must="M,V" may-encrypt="N" vendor-id="10415">
            <!-- http://www.qtc.jp/3GPP/Specs/29273-920.pdf Section 8.2.3.9-->
            <data type="Grouped">
                <rule avp="SIP-Item-Number" required="false" max="1"/>
                <rule avp="SIP-Authentication-Scheme" required="false" max="1"/>
                <rule avp="SIP-Authenticate" required="false" max="1"/>
                <rule avp="SIP-Authorization" required="false" max="1"/>
                <rule avp="Confidentiality-Key" required="false" max="1"/>
                <rule avp="Integrity-Key" required="false" max="1"/>
                <rule avp="AVP" required="false"/>
            </data>
        </avp>

        <avp name="SIP-Item-Number" code="613" must="M,V" may-encrypt="N" vendor-id="10415">
            <!-- 3GPP TS 29.229 Section 6.3.14 -->
            <data type="Unsigned32"/>
        </avp>

        <avp name="SIP-Authentication-Scheme" code="608" must="M,V" may-encrypt="N" vendor-id="10415">
             <!-- 3GPP TS 29.229 Section 6.3.9 -->
            <data type="UTF8String"/>
        </avp>

        <avp name="SIP-Authenticate" code="609" must="M,V" may-encrypt="N" vendor-id="10415">
            <!-- 3GPP TS 29.229 Section 6.3.10 -->
            <data type="OctetString"/>
        </avp>

        <avp name="SIP-Authorization" code="610" must="M,V" may-encrypt="N" vendor-id="10415">
            <!-- 3GPP TS 29.229 Section 6.3.11 -->
            <data type="OctetString"/>
        </avp>

        <avp name="Confidentiality-Key" code="625" must="M,V" may-encrypt="N" vendor-id="10415">
            <!-- http://www.qtc.jp/3GPP/Specs/29273-920.pdf Section 8.2.3.10 -->
            <data type="OctetString"/>
        </avp>

        <avp name="Integrity-Key" code="626" must="M,V" may-encrypt="N" vendor-id="10415">
            <!-- http://www.qtc.jp/3GPP/Specs/29273-920.pdf Section 8.2.3.11 -->
            <data type="OctetString"/>
        </avp>

        <avp name="SIP-Number-Auth-Items" code="607" must="M,V" may-encrypt="N" vendor-id="10415">
            <!-- 3GPP TS 29.229 Section 6.3.8 -->
            <data type="Unsigned32" />
        </avp>

        <avp name="TGPP-AAA-Server-Name" code="318" must="M,V" may-encrypt="N" vendor-id="10415">
            <!-- 3GPP TS 29.234 Section 10.1.34 -->
            <data type="DiameterIdentity"/>
        </avp>

        <avp name="Supported-Features" code="628" vendor-id="10415" must="V" may="M" may-encrypt="N">
            <!-- 3GPP TS 29.229 Section 6.3.29 -->
            <data type="Grouped">
                <rule avp="Vendor-Id" required="true" max="1"/>
                <rule avp="Feature-List-ID" required="true" max="1"/>
                <rule avp="Feature-List" required="true" max="1"/>
            </data>
        </avp>

        <avp name="Service-Selection" code="493" must="M" may="P" must-not="V" may-encrypt="Y" vendor-id="0">
            <!-- http://www.qtc.jp/3GPP/Specs/29273-920.pdf Section 5.2.3.5 -->
            <data type="UTF8String"/>
        </avp>

        <avp name="Context-Identifier" code="1423" must="M,V" may-encrypt="N" vendor-id="10415">
            <!-- 3GPP TS 29.272 Section 7.3.27 -->
            <data type="Unsigned32"/>
        </avp>

        <!-- RFC 5447 Diameter Mobile IPv6: Support for Network Access Server to Diameter Server Interaction -->
        <avp name="MIP6-Agent-Info" code="486" must="M" may="P" must-not="V" may-encrypt="Y" vendor-id="10415">
            <data type="Grouped">
                <rule avp="MIP-Home-Agent-Address" required="false" max="2"/>
                <rule avp="MIP-Home-Agent-Host" required="false" max="1"/>
                <rule avp="MIP6-Home-Link-Prefix" required="false" max="1"/>
                <rule avp="AVP" required="false"/>
            </data>
        </avp>

        <avp name="Server-Assignment-Type" code="614" must="M,V" may-encrypt="N" vendor-id="10415">
            <!-- http://www.qtc.jp/3GPP/Specs/29273-920.pdf Section 8.2.3.12 -->
            <data type="Enumerated">
                <item code="0" name="NO_ASSIGNMENT"/>
                <item code="1" name="REGISTRATION"/>
                <item code="2" name="RE_REGISTRATION"/>
                <item code="3" name="UNREGISTERED_USER"/>
                <item code="4" name="TIMEOUT_DEREGISTRATION"/>
                <item code="5" name="USER_DEREGISTRATION"/>
                <item code="6" name="TIMEOUT_DEREGISTRATION_STORE_SERVER_NAME"/>
                <item code="7" name="USER_DEREGISTRATION_STORE_SERVER_NAME"/>
                <item code="8" name="ADMINISTRATIVE_DEREGISTRATION"/>
                <item code="9" name="AUTHENTICATION_FAILURE"/>
                <item code="10" name="AUTHENTICATION_TIMEOUT"/>
                <item code="11" name="DEREGISTRATION_TOO_MUCH_DATA"/>
                <item code="12" name="AAA_USER_DATA_REQUEST"/>
                <item code="13" name="PGW_UPDATE"/>
                <item code="14" name="RESTORATION"/>
            </data>
        </avp>

        <avp name="Deregistration-Reason" code="615" must="M,V" may-encrypt="N" vendor-id="10415">
             <!-- https://www.etsi.org/deliver/etsi_ts/129200_129299/129229/10.05.00_60/ts_129229v100500p.pdf -->
             <data type="Grouped">
                <rule avp="Reason-Code" required="true"/>
                <rule avp="Reason-Info" required="false"/>
                <rule avp="AVP" required="false"/>
             </data>
        </avp>

        <avp name="Reason-Code" code="616" must="M,V" may-encrypt="N" vendor-id="10415">
             <!-- https://www.etsi.org/deliver/etsi_ts/129200_129299/129229/10.05.00_60/ts_129229v100500p.pdf -->
             <data type="Enumerated">
                <item code="0" name="PERMANENT_TERMINATION"/>
                <item code="1" name="NEW_SERVER_ASSIGNMENT"/>
                <item code="2" name="SERVER_CHANGE"/>
                <item code="3" name="REMOVE_S_CSCF"/>
             </data>
        </avp>

        <avp name="Reason-Info" code="617" must="M,V" may-encrypt="N" vendor-id="10415">
             <!-- https://www.etsi.org/deliver/etsi_ts/129200_129299/129229/10.05.00_60/ts_129229v100500p.pdf -->
             <data type="UTF8String"/>
        </avp>

        <avp name="Non-3GPP-User-Data" code="1500" must="M,V" may-encrypt="N" vendor-id="10415">
            <!-- http://www.qtc.jp/3GPP/Specs/29273-920.pdf Section 8.2.3.1 -->
            <data type="Grouped">
                <rule avp="Subscription-Id" required="false" max="1"/>
                <rule avp="Non-3GPP-IP-Access" required="false" max="1"/>
                <rule avp="Non-3GPP-IP-Access-APN" required="false" max="1"/>
                <rule avp="RAT-Type" required="false"/>
                <rule avp="Session-Timeout" required="false" max="1"/>
                <rule avp="MIP6-Feature-Vector" required="false" max="1"/>
                <rule avp="AMBR" required="false" max="1"/>
                <rule avp="3GPP-Charging-Characteristics" required="false" max="1"/>
                <rule avp="Context-Identifier" required="false" max="1"/>
                <rule avp="APN-OI-Replacement" required="false" max="1"/>
                <rule avp="APN-Configuration" required="false"/>
                <rule avp="Trace-Info" required="false" max="1"/>
                <rule avp="AVP" required="false"/>
            </data>
        </avp>

        <avp name="Subscription-Id" code="443" must="M" may="P" must-not="V" may-encrypt="Y" vendor-id="0">
            <!-- https://tools.ietf.org/rfc/rfc4006.txt -->
            <data type="Grouped">
                <rule avp="Subscription-Id-Type" required="false" max="1"/>
                <rule avp="Subscription-Id-Data" required="false" max="1"/>
            </data>
        </avp>

        <avp name="Subscription-Id-Type" code="450" must="M" may="P" must-not="V" may-encrypt="Y" vendor-id="0">
            <!-- https://tools.ietf.org/rfc/rfc4006.txt -->
            <data type="Enumerated">
                <item code="0" name="END_USER_E164"/>
                <item code="1" name="END_USER_IMSI"/>
                <item code="2" name="END_USER_SIP_URI"/>
                <item code="3" name="END_USER_NAI"/>
                <item code="4" name="END_USER_PRIVATE"/>
            </data>
        </avp>

        <avp name="Subscription-Id-Data" code="444" must="M" may="P" must-not="V" may-encrypt="Y" vendor-id="0">
            <!-- https://tools.ietf.org/rfc/rfc4006.txt -->
            <data type="UTF8String"/>
        </avp>

        <avp name="Non-3GPP-IP-Access" code="1501" must="M,V" may-encrypt="N" vendor-id="10415">
            <!-- http://www.qtc.jp/3GPP/Specs/29273-920.pdf Section 8.2.3.3 -->
            <data type="Enumerated">
                <item code="0" name="NON_3GPP_SUBSCRIPTION_ALLOWED"/>
                <item code="1" name="NON_3GPP_SUBSCRIPTION_BARRED"/>
            </data>
        </avp>

        <avp name="Non-3GPP-IP-Access-APN" code="1502" must="M,V" may-encrypt="N" vendor-id="10415">
            <!-- http://www.qtc.jp/3GPP/Specs/29273-920.pdf Section 8.2.3.4 -->
            <data type="Enumerated">
                <item code="0" name="NON_3GPP_APNS_ENABLE"/>
                <item code="1" name="NON_3GPP_APNS_DISABLE"/>
            </data>
        </avp>

        <avp name="MIP6-Feature-Vector" code="124" must="M" may="P" may-encrypt="N" vendor-id="0">
            <!-- http://www.qtc.jp/3GPP/Specs/29273-920.pdf Section 5.2.3.3 -->
            <data type="Unsigned64"/>
        </avp>

        <avp name="AMBR" code="1435" must="M,V" may-encrypt="N" vendor-id="10415">
            <!-- 3GPP TS 29.272 Section 7.3.41 -->
            <data type="Grouped">
                <rule avp="Max-Requested-Bandwidth-UL" required="true" max="1"/>
                <rule avp="Max-Requested-Bandwidth-DL" required="true" max="1"/>
                <rule avp="AVP" required="false"/>
            </data>
        </avp>

        <avp name="Max-Requested-Bandwidth-DL" code="515" must="V,M" may="P" must-not="-" may-encrypt="Y" vendor-id="10415">
            <!-- 3GPP TS 29.214 [11] -->
            <data type="Unsigned32"/>
        </avp>

        <avp name="Max-Requested-Bandwidth-UL" code="516" must="V,M" may="P" must-not="-" may-encrypt="Y" vendor-id="10415">
            <!-- 3GPP TS 29.214 [11] -->
            <data type="Unsigned32"/>
        </avp>

        <avp name="TGPP-Charging-Characteristics" code="13" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
            <!-- 3GPP TS 29.061 [21] -->
            <data type="UTF8String"/>
        </avp>

        <avp name="APN-OI-Replacement" code="1427" must="M,V" may-encrypt="N" vendor-id="10415">
            <!-- 3GPP TS 29.272 Section 7.3.32 -->
            <data type="UTF8String"/>
        </avp>

        <avp name="APN-Configuration" code="1430" must="M,V" may-encrypt="N" vendor-id="10415">
            <!-- 3GPP TS 29.272 Section 7.3.35 -->
            <data type="Grouped">
                <rule avp="Context-Identifier" required="true" max="1"/>
                <rule avp="Served-Party-IP-Address" required="false" max="2"/>
                <rule avp="PDN-Type" required="true" max="1"/>
                <rule avp="Service-Selection" required="true" max="1"/>
                <rule avp="EPS-Subscribed-QoS-Profile" required="false" max="1"/>
                <rule avp="VPLMN-Dynamic-Address-Allowed" required="false" max="1"/>
                <rule avp="MIP6-Agent-Info" required="false" max="1"/>
                <rule avp="Visited-Network-Identifier" required="false" max="1"/>
                <rule avp="PDN-GW-Allocation-Type" required="false" max="1"/>
                <rule avp="TGPP-Charging-Characteristics" required="false" max="1"/>
                <rule avp="AMBR" required="false" max="1"/>
                <rule avp="Specific-APN-Info" required="false"/>
                <rule avp="APN-OI-Replacement" required="false" max="1"/>
                <rule avp="SIPTO-Permission" required="false" max="1"/>
                <rule avp="LIPA-Permission" required="false" max="1"/>
                <rule avp="AVP" required="false"/>
            </data>
        </avp>

        <avp name="Served-Party-IP-Address" code="848" must="M,V" may="P" may-encrypt="N" vendor-id="10415">
            <!-- 3GPP TS 32.299 [8] -->
            <data type="Address"/>
        </avp>

        <avp name="PDN-Type" code="1456" must="M,V" may-encrypt="N" vendor-id="10415">
            <!-- 3GPP TS 29.272 Section 7.3.62 -->
            <data type="Enumerated">
                <item code="0" name="IPv4"/>
                <item code="1" name="IPv6"/>
                <item code="2" name="IPv4v6"/>
                <item code="3" name="IPv4_OR_IPv6"/>
            </data>
        </avp>

        <avp name="EPS-Subscribed-QoS-Profile" code="1431" must="M,V" may-encrypt="N" vendor-id="10415">
            <!-- 3GPP TS 29.272 Section 7.3.37 -->
            <data type="Grouped">
                <rule avp="QoS-Class-Identifier" required="true" max="1"/>
                <rule avp="Allocation-Retention-Priority" required="true" max="1"/>
                <rule avp="AVP" required="false"/>
            </data>
        </avp>

        <avp name="QoS-Class-Identifier" code="1028" must="V,M" may="P" must-not="-" may-encrypt="Y" vendor-id="10415">
            <!-- 3GPP TS 29.212 [10] -->
            <data type="Enumerated">
                <item code="1" name="QCI_1"/>
                <item code="2" name="QCI_2"/>
                <item code="3" name="QCI_3"/>
                <item code="4" name="QCI_4"/>
                <item code="5" name="QCI_5"/>
                <item code="6" name="QCI_6"/>
                <item code="7" name="QCI_7"/>
                <item code="8" name="QCI_8"/>
                <item code="9" name="QCI_9"/>
                <item code="65" name="QCI_65"/>
                <item code="66" name="QCI_66"/>
                <item code="69" name="QCI_69"/>
                <item code="70" name="QCI_70"/>
                <item code="75" name="QCI_75"/>
                <item code="79" name="QCI_79"/>
            </data>
        </avp>

        <avp name="Allocation-Retention-Priority" code="1034" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
            <!-- 3GPP TS 29.212 [10] -->
            <data type="Grouped">
                <rule avp="Priority-Level" required="true" max="1"/>
                <rule avp="Pre-emption-Capability" required="false" max="1"/>
                <rule avp="Pre-emption-Vulnerability" required="false" max="1"/>
            </data>
        </avp>

        <avp name="Priority-Level" code="1046" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
            <!-- 3GPP TS 29.212 [10] -->
            <data type="Unsigned32"/>
        </avp>

        <avp name="Pre-emption-Capability" code="1047" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
            <!-- 3GPP TS 29.212 [10] -->
            <data type="Enumerated">
                <item code="0" name="PRE-EMPTION_CAPABILITY_ENABLED"/>
                <item code="1" name="PRE-EMPTION_CAPABILITY_DISABLED"/>
            </data>
        </avp>

        <avp name="Pre-emption-Vulnerability" code="1048" must="V" may="P" must-not="M" may-encrypt="Y" vendor-id="10415">
            <!-- 3GPP TS 29.212 [10] -->
            <data type="Enumerated">
                <item code="0" name="PRE-EMPTION_VULNERABILITY_ENABLED"/>
                <item code="1" name="PRE-EMPTION_VULNERABILITY_DISABLED"/>
            </data>
        </avp>

        <avp name="VPLMN-Dynamic-Address-Allowed" code="1432" must="M,V" may-encrypt="N" vendor-id="10415">
            <!-- 3GPP TS 29.272 Section 7.3.38 -->
            <data type="Enumerated">
                <item code="0" name="NOTALLOWED"/>
                <item code="1" name="ALLOWED"/>
            </data>
        </avp>

        <avp name="PDN-GW-Allocation-Type" code="1438" must="M,V" may-encrypt="N" vendor-id="10415">
            <!-- 3GPP TS 29.272 Section 7.3.44 -->
            <data type="Enumerated">
                <item code="0" name="STATIC"/>
                <item code="1" name="DYNAMIC"/>
            </data>
        </avp>

        <avp name="Specific-APN-Info" code="1472" vendor-id="10415" must="M,V" may-encrypt="N">
            <!-- 3GPP TS 29.272 Section 7.3.82 -->
            <data type="Grouped">
                <rule avp="Service-Selection" required="true" max="1"/>
                <rule avp="MIP6-Agent-Info" required="true" max="1"/>
                <rule avp="Visited-Network-Identifier" required="false" max="1"/>
                <rule avp="AVP" required="false"/>
            </data>
        </avp>

        <avp name="SIPTO-Permission" code="1613" must="V" must-not="M" may-encrypt="N" vendor-id="10415">
            <!-- 3GPP TS 29.272 Section 7.3.135 -->
            <data type="Enumerated">
                <item code="0" name="SIPTO_ALLOWED"/>
                <item code="1" name="SIPTO_NOTALLOWED"/>
            </data>
        </avp>

        <avp name="LIPA-Permission" code="1618" must="V" must-not="M" may-encrypt="N" vendor-id="10415">
            <!-- 3GPP TS 29.272 Section 7.3.132 -->
            <data type="Enumerated">
                <item code="0" name="LIPA-PROHIBITED"/>
                <item code="1" name="LIPA-ONLY"/>
                <item code="2" name="LIPA-CONDITIONAL"/>
            </data>
        </avp>

        <avp name="Trace-Info" code="1505" must="V" must-not="M" vendor-id="10415">
            <!-- http://www.qtc.jp/3GPP/Specs/29273-920.pdf Section 8.2.3.13 -->
            <data type="Grouped">
                <rule avp="Trace-Data" required="false" max="1"/>
                <rule avp="Trace-Reference" required="false" max="1"/>
            </data>
        </avp>

        <avp name="Trace-Data" code="1458" must="M,V" may-encrypt="N" vendor-id="10415">
            <!-- 3GPP TS 29.272 Section 7.3.63 -->
            <data type="Grouped">
                <rule avp="Trace-Reference" required="true" max="1"/>
                <rule avp="Trace-Depth" required="true" max="1"/>
                <rule avp="Trace-NE-Type-List" required="true" max="1"/>
                <rule avp="Trace-Interface-List" required="false" max="1"/>
                <rule avp="Trace-Event-List" required="true" max="1"/>
                <rule avp="OMC-Id" required="false" max="1"/>
                <rule avp="Trace-Collection-Entity" required="true" max="1"/>
                <rule avp="MDT-Configuration" required="false" max="1"/>
                <rule avp="AVP" required="false"/>
            </data>
        </avp>

        <avp name="Trace-Reference" code="1459" must="M,V" may-encrypt="N" vendor-id="10415">
            <!-- 3GPP TS 29.272 Section 7.3.64 -->
            <data type="OctetString"/>
        </avp>

        <avp name="Trace-Depth" code="1462" must="M,V" may-encrypt="N" vendor-id="10415">
            <!-- 3GPP TS 29.272 Section 7.3.67 -->
            <data type="Enumerated">
                <item code="0" name="LIPA-PROHIBITED"/>
                <item code="1" name="LIPA-ONLY"/>
                <item code="2" name="LIPA-CONDITIONAL"/>
            </data>
        </avp>

        <avp name="Trace-NE-Type-List" code="1463" must="M,V" may-encrypt="N" vendor-id="10415">
            <!-- 3GPP TS 29.272 Section 7.3.68 -->
            <data type="OctetString"/>
        </avp>

        <avp name="Trace-Interface-List" code="1464" must="M,V" may-encrypt="N" vendor-id="10415">
            <!-- 3GPP TS 29.272 Section 7.3.69 -->
            <data type="OctetString"/>
        </avp>

        <avp name="Trace-Event-List" code="1465" must="M,V" may-encrypt="N" vendor-id="10415">
            <!-- 3GPP TS 29.272 Section 7.3.70 -->
            <data type="OctetString"/>
        </avp>

        <avp name="OMC-Id" code="1466" must="M,V" may-encrypt="N" vendor-id="10415">
            <!-- 3GPP TS 29.272 Section 7.3.71 -->
            <data type="OctetString"/>
        </avp>

        <avp name="Trace-Collection-Entity" code="1452" must="M,V" may="P" may-encrypt="N" vendor-id="10415">
            <!-- 3GPP TS 29.272 Section 7.3.98 -->
            <data type="Address"/>
        </avp>

        <avp name="MDT-Configuration" code="1622" must="M,V" may-encrypt="N" vendor-id="10415">
            <!-- 3GPP TS 29.272 Section 7.3.136 -->
            <data type="Grouped">
                <rule avp="QoS-Class-Identifier" required="true" max="1"/>
                <rule avp="Allocation-Retention-Priority" required="true" max="1"/>
                <rule avp="AVP" required="false"/>
            </data>
        </avp>

    </application>
</diameter>`,
}
