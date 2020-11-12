package library

var DiameterSy = DictInfo{
	Name: "Diameter Sy",
	XML: `<?xml version="1.0" encoding="UTF-8"?>
<diameter>

	<application id="16777302" type="auth" name="Diameter Sy">
		<!-- Diameter Credit Control Application -->
		<!-- http://tools.ietf.org/html/rfc4006 -->

		<command code="8388635" short="SL" name="Spending-Limit">
			<request>
				<!-- http://tools.ietf.org/html/rfc4006#section-3.1 -->
				<rule avp="Session-Id" required="true" max="1"/>
				<rule avp="Auth-Application-Id" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="Destination-Realm" required="true" max="1"/>
				<rule avp="SL-Request-Type" required="true" max="1"/>
				<rule avp="Destination-Host" required="false" max="1"/>
				<rule avp="Origin-State-Id" required="false" max="1"/>
				<rule avp="Subscription-Id" required="false" max="1"/>
				<rule avp="Policy-Counter-Identifier" required="false" max="1"/>
				<rule avp="Proxy-Info" required="false" max="1"/>
				<rule avp="Route-Record" required="false" max="1"/>
				<rule avp="Service-Information" required="false" max="1"/>
			</request>
			<answer>
				<!-- http://tools.ietf.org/html/rfc4006#section-3.2 -->
				<rule avp="Session-Id" required="true" max="1"/>
				<rule avp="Result-Code" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="Origin-State-Id" required="false" max="1"/>
				<rule avp="Redirect-Host" required="false" max="1"/>
				<rule avp="Redirect-Host-Usage" required="false" max="1"/>
				<rule avp="Redirect-Max-Cache-Time" required="false" max="1"/>
				<rule avp="Proxy-Info" required="false" max="1"/>
				<rule avp="Route-Record" required="false" max="1"/>
				<rule avp="Failed-AVP" required="false" max="1"/>
			</answer>
		</command>

		<avp name="SL-Request-Type" code="2904" must="M" may="P" must-not="V" may-encrypt="-">
			<data type="Enumerated">
				<item code="0" name="INITIAL_REQUEST"/>
				<item code="1" name="INTERMEDIATE_REQUEST"/>
			</data>
		</avp>

    </application>
</diameter>`,
}
