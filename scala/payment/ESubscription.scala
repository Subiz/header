// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO2

package payment

/** subscription with error
  */
@SerialVersionUID(0L)
final case class ESubscription(
    sub: scala.Option[payment.Subscription] = None,
    err: scala.Option[_root_.scala.Predef.String] = None
    ) extends scalapb.GeneratedMessage with scalapb.Message[ESubscription] with scalapb.lenses.Updatable[ESubscription] {
    @transient
    private[this] var __serializedSizeCachedValue: _root_.scala.Int = 0
    private[this] def __computeSerializedValue(): _root_.scala.Int = {
      var __size = 0
      if (sub.isDefined) { __size += 1 + _root_.com.google.protobuf.CodedOutputStream.computeUInt32SizeNoTag(sub.get.serializedSize) + sub.get.serializedSize }
      if (err.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(3, err.get) }
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
      sub.foreach { __v =>
        _output__.writeTag(2, 2)
        _output__.writeUInt32NoTag(__v.serializedSize)
        __v.writeTo(_output__)
      };
      err.foreach { __v =>
        _output__.writeString(3, __v)
      };
    }
    def mergeFrom(`_input__`: _root_.com.google.protobuf.CodedInputStream): payment.ESubscription = {
      var __sub = this.sub
      var __err = this.err
      var _done__ = false
      while (!_done__) {
        val _tag__ = _input__.readTag()
        _tag__ match {
          case 0 => _done__ = true
          case 18 =>
            __sub = Option(_root_.scalapb.LiteParser.readMessage(_input__, __sub.getOrElse(payment.Subscription.defaultInstance)))
          case 26 =>
            __err = Option(_input__.readString())
          case tag => _input__.skipField(tag)
        }
      }
      payment.ESubscription(
          sub = __sub,
          err = __err
      )
    }
    def getSub: payment.Subscription = sub.getOrElse(payment.Subscription.defaultInstance)
    def clearSub: ESubscription = copy(sub = None)
    def withSub(__v: payment.Subscription): ESubscription = copy(sub = Option(__v))
    def getErr: _root_.scala.Predef.String = err.getOrElse("")
    def clearErr: ESubscription = copy(err = None)
    def withErr(__v: _root_.scala.Predef.String): ESubscription = copy(err = Option(__v))
    def getFieldByNumber(__fieldNumber: _root_.scala.Int): scala.Any = {
      (__fieldNumber: @_root_.scala.unchecked) match {
        case 2 => sub.orNull
        case 3 => err.orNull
      }
    }
    def getField(__field: _root_.scalapb.descriptors.FieldDescriptor): _root_.scalapb.descriptors.PValue = {
      require(__field.containingMessage eq companion.scalaDescriptor)
      (__field.number: @_root_.scala.unchecked) match {
        case 2 => sub.map(_.toPMessage).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 3 => err.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
      }
    }
    def toProtoString: _root_.scala.Predef.String = _root_.scalapb.TextFormat.printToUnicodeString(this)
    def companion = payment.ESubscription
}

object ESubscription extends scalapb.GeneratedMessageCompanion[payment.ESubscription] {
  implicit def messageCompanion: scalapb.GeneratedMessageCompanion[payment.ESubscription] = this
  def fromFieldsMap(__fieldsMap: scala.collection.immutable.Map[_root_.com.google.protobuf.Descriptors.FieldDescriptor, scala.Any]): payment.ESubscription = {
    require(__fieldsMap.keys.forall(_.getContainingType() == javaDescriptor), "FieldDescriptor does not match message type.")
    val __fields = javaDescriptor.getFields
    payment.ESubscription(
      __fieldsMap.get(__fields.get(0)).asInstanceOf[scala.Option[payment.Subscription]],
      __fieldsMap.get(__fields.get(1)).asInstanceOf[scala.Option[_root_.scala.Predef.String]]
    )
  }
  implicit def messageReads: _root_.scalapb.descriptors.Reads[payment.ESubscription] = _root_.scalapb.descriptors.Reads{
    case _root_.scalapb.descriptors.PMessage(__fieldsMap) =>
      require(__fieldsMap.keys.forall(_.containingMessage == scalaDescriptor), "FieldDescriptor does not match message type.")
      payment.ESubscription(
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(2).get).flatMap(_.as[scala.Option[payment.Subscription]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(3).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]])
      )
    case _ => throw new RuntimeException("Expected PMessage")
  }
  def javaDescriptor: _root_.com.google.protobuf.Descriptors.Descriptor = PaymentProto.javaDescriptor.getMessageTypes.get(29)
  def scalaDescriptor: _root_.scalapb.descriptors.Descriptor = PaymentProto.scalaDescriptor.messages(29)
  def messageCompanionForFieldNumber(__number: _root_.scala.Int): _root_.scalapb.GeneratedMessageCompanion[_] = {
    var __out: _root_.scalapb.GeneratedMessageCompanion[_] = null
    (__number: @_root_.scala.unchecked) match {
      case 2 => __out = payment.Subscription
    }
    __out
  }
  lazy val nestedMessagesCompanions: Seq[_root_.scalapb.GeneratedMessageCompanion[_]] = Seq.empty
  def enumCompanionForFieldNumber(__fieldNumber: _root_.scala.Int): _root_.scalapb.GeneratedEnumCompanion[_] = throw new MatchError(__fieldNumber)
  lazy val defaultInstance = payment.ESubscription(
  )
  implicit class ESubscriptionLens[UpperPB](_l: _root_.scalapb.lenses.Lens[UpperPB, payment.ESubscription]) extends _root_.scalapb.lenses.ObjectLens[UpperPB, payment.ESubscription](_l) {
    def sub: _root_.scalapb.lenses.Lens[UpperPB, payment.Subscription] = field(_.getSub)((c_, f_) => c_.copy(sub = Option(f_)))
    def optionalSub: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[payment.Subscription]] = field(_.sub)((c_, f_) => c_.copy(sub = f_))
    def err: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getErr)((c_, f_) => c_.copy(err = Option(f_)))
    def optionalErr: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.err)((c_, f_) => c_.copy(err = f_))
  }
  final val SUB_FIELD_NUMBER = 2
  final val ERR_FIELD_NUMBER = 3
}