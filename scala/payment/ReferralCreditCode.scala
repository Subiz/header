// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO2

package payment

@SerialVersionUID(0L)
final case class ReferralCreditCode(
    referrerId: scala.Option[_root_.scala.Predef.String] = None,
    credit: scala.Option[_root_.scala.Float] = None
    ) extends scalapb.GeneratedMessage with scalapb.Message[ReferralCreditCode] with scalapb.lenses.Updatable[ReferralCreditCode] {
    @transient
    private[this] var __serializedSizeCachedValue: _root_.scala.Int = 0
    private[this] def __computeSerializedValue(): _root_.scala.Int = {
      var __size = 0
      if (referrerId.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(2, referrerId.get) }
      if (credit.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeFloatSize(3, credit.get) }
      __size
    }
    final override def serializedSize: _root_.scala.Int = {
      var read = __serializedSizeCachedValue
      if (read == 0) {
        read = __computeSerializedValue()
        __serializedSizeCachedValue = read
      }
      read
    }
    def writeTo(`_output__`: _root_.com.google.protobuf.CodedOutputStream): _root_.scala.Unit = {
      referrerId.foreach { __v =>
        _output__.writeString(2, __v)
      };
      credit.foreach { __v =>
        _output__.writeFloat(3, __v)
      };
    }
    def mergeFrom(`_input__`: _root_.com.google.protobuf.CodedInputStream): payment.ReferralCreditCode = {
      var __referrerId = this.referrerId
      var __credit = this.credit
      var _done__ = false
      while (!_done__) {
        val _tag__ = _input__.readTag()
        _tag__ match {
          case 0 => _done__ = true
          case 18 =>
            __referrerId = Option(_input__.readString())
          case 29 =>
            __credit = Option(_input__.readFloat())
          case tag => _input__.skipField(tag)
        }
      }
      payment.ReferralCreditCode(
          referrerId = __referrerId,
          credit = __credit
      )
    }
    def getReferrerId: _root_.scala.Predef.String = referrerId.getOrElse("")
    def clearReferrerId: ReferralCreditCode = copy(referrerId = None)
    def withReferrerId(__v: _root_.scala.Predef.String): ReferralCreditCode = copy(referrerId = Option(__v))
    def getCredit: _root_.scala.Float = credit.getOrElse(0.0f)
    def clearCredit: ReferralCreditCode = copy(credit = None)
    def withCredit(__v: _root_.scala.Float): ReferralCreditCode = copy(credit = Option(__v))
    def getFieldByNumber(__fieldNumber: _root_.scala.Int): scala.Any = {
      (__fieldNumber: @_root_.scala.unchecked) match {
        case 2 => referrerId.orNull
        case 3 => credit.orNull
      }
    }
    def getField(__field: _root_.scalapb.descriptors.FieldDescriptor): _root_.scalapb.descriptors.PValue = {
      require(__field.containingMessage eq companion.scalaDescriptor)
      (__field.number: @_root_.scala.unchecked) match {
        case 2 => referrerId.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 3 => credit.map(_root_.scalapb.descriptors.PFloat).getOrElse(_root_.scalapb.descriptors.PEmpty)
      }
    }
    def toProtoString: _root_.scala.Predef.String = _root_.scalapb.TextFormat.printToUnicodeString(this)
    def companion = payment.ReferralCreditCode
}

object ReferralCreditCode extends scalapb.GeneratedMessageCompanion[payment.ReferralCreditCode] {
  implicit def messageCompanion: scalapb.GeneratedMessageCompanion[payment.ReferralCreditCode] = this
  def fromFieldsMap(__fieldsMap: scala.collection.immutable.Map[_root_.com.google.protobuf.Descriptors.FieldDescriptor, scala.Any]): payment.ReferralCreditCode = {
    require(__fieldsMap.keys.forall(_.getContainingType() == javaDescriptor), "FieldDescriptor does not match message type.")
    val __fields = javaDescriptor.getFields
    payment.ReferralCreditCode(
      __fieldsMap.get(__fields.get(0)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(1)).asInstanceOf[scala.Option[_root_.scala.Float]]
    )
  }
  implicit def messageReads: _root_.scalapb.descriptors.Reads[payment.ReferralCreditCode] = _root_.scalapb.descriptors.Reads{
    case _root_.scalapb.descriptors.PMessage(__fieldsMap) =>
      require(__fieldsMap.keys.forall(_.containingMessage == scalaDescriptor), "FieldDescriptor does not match message type.")
      payment.ReferralCreditCode(
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(2).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(3).get).flatMap(_.as[scala.Option[_root_.scala.Float]])
      )
    case _ => throw new RuntimeException("Expected PMessage")
  }
  def javaDescriptor: _root_.com.google.protobuf.Descriptors.Descriptor = PaymentProto.javaDescriptor.getMessageTypes.get(21)
  def scalaDescriptor: _root_.scalapb.descriptors.Descriptor = PaymentProto.scalaDescriptor.messages(21)
  def messageCompanionForFieldNumber(__number: _root_.scala.Int): _root_.scalapb.GeneratedMessageCompanion[_] = throw new MatchError(__number)
  lazy val nestedMessagesCompanions: Seq[_root_.scalapb.GeneratedMessageCompanion[_]] = Seq.empty
  def enumCompanionForFieldNumber(__fieldNumber: _root_.scala.Int): _root_.scalapb.GeneratedEnumCompanion[_] = throw new MatchError(__fieldNumber)
  lazy val defaultInstance = payment.ReferralCreditCode(
  )
  implicit class ReferralCreditCodeLens[UpperPB](_l: _root_.scalapb.lenses.Lens[UpperPB, payment.ReferralCreditCode]) extends _root_.scalapb.lenses.ObjectLens[UpperPB, payment.ReferralCreditCode](_l) {
    def referrerId: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getReferrerId)((c_, f_) => c_.copy(referrerId = Option(f_)))
    def optionalReferrerId: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.referrerId)((c_, f_) => c_.copy(referrerId = f_))
    def credit: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Float] = field(_.getCredit)((c_, f_) => c_.copy(credit = Option(f_)))
    def optionalCredit: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Float]] = field(_.credit)((c_, f_) => c_.copy(credit = f_))
  }
  final val REFERRER_ID_FIELD_NUMBER = 2
  final val CREDIT_FIELD_NUMBER = 3
}