// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO2

package payment

@SerialVersionUID(0L)
final case class PaymentMethods(
    paymentMethods: _root_.scala.collection.Seq[payment.PaymentMethod] = _root_.scala.collection.Seq.empty
    ) extends scalapb.GeneratedMessage with scalapb.Message[PaymentMethods] with scalapb.lenses.Updatable[PaymentMethods] {
    @transient
    private[this] var __serializedSizeCachedValue: _root_.scala.Int = 0
    private[this] def __computeSerializedValue(): _root_.scala.Int = {
      var __size = 0
      paymentMethods.foreach(paymentMethods => __size += 1 + _root_.com.google.protobuf.CodedOutputStream.computeUInt32SizeNoTag(paymentMethods.serializedSize) + paymentMethods.serializedSize)
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
      paymentMethods.foreach { __v =>
        _output__.writeTag(2, 2)
        _output__.writeUInt32NoTag(__v.serializedSize)
        __v.writeTo(_output__)
      };
    }
    def mergeFrom(`_input__`: _root_.com.google.protobuf.CodedInputStream): payment.PaymentMethods = {
      val __paymentMethods = (_root_.scala.collection.immutable.Vector.newBuilder[payment.PaymentMethod] ++= this.paymentMethods)
      var _done__ = false
      while (!_done__) {
        val _tag__ = _input__.readTag()
        _tag__ match {
          case 0 => _done__ = true
          case 18 =>
            __paymentMethods += _root_.scalapb.LiteParser.readMessage(_input__, payment.PaymentMethod.defaultInstance)
          case tag => _input__.skipField(tag)
        }
      }
      payment.PaymentMethods(
          paymentMethods = __paymentMethods.result()
      )
    }
    def clearPaymentMethods = copy(paymentMethods = _root_.scala.collection.Seq.empty)
    def addPaymentMethods(__vs: payment.PaymentMethod*): PaymentMethods = addAllPaymentMethods(__vs)
    def addAllPaymentMethods(__vs: TraversableOnce[payment.PaymentMethod]): PaymentMethods = copy(paymentMethods = paymentMethods ++ __vs)
    def withPaymentMethods(__v: _root_.scala.collection.Seq[payment.PaymentMethod]): PaymentMethods = copy(paymentMethods = __v)
    def getFieldByNumber(__fieldNumber: _root_.scala.Int): scala.Any = {
      (__fieldNumber: @_root_.scala.unchecked) match {
        case 2 => paymentMethods
      }
    }
    def getField(__field: _root_.scalapb.descriptors.FieldDescriptor): _root_.scalapb.descriptors.PValue = {
      require(__field.containingMessage eq companion.scalaDescriptor)
      (__field.number: @_root_.scala.unchecked) match {
        case 2 => _root_.scalapb.descriptors.PRepeated(paymentMethods.map(_.toPMessage)(_root_.scala.collection.breakOut))
      }
    }
    def toProtoString: _root_.scala.Predef.String = _root_.scalapb.TextFormat.printToUnicodeString(this)
    def companion = payment.PaymentMethods
}

object PaymentMethods extends scalapb.GeneratedMessageCompanion[payment.PaymentMethods] {
  implicit def messageCompanion: scalapb.GeneratedMessageCompanion[payment.PaymentMethods] = this
  def fromFieldsMap(__fieldsMap: scala.collection.immutable.Map[_root_.com.google.protobuf.Descriptors.FieldDescriptor, scala.Any]): payment.PaymentMethods = {
    require(__fieldsMap.keys.forall(_.getContainingType() == javaDescriptor), "FieldDescriptor does not match message type.")
    val __fields = javaDescriptor.getFields
    payment.PaymentMethods(
      __fieldsMap.getOrElse(__fields.get(0), Nil).asInstanceOf[_root_.scala.collection.Seq[payment.PaymentMethod]]
    )
  }
  implicit def messageReads: _root_.scalapb.descriptors.Reads[payment.PaymentMethods] = _root_.scalapb.descriptors.Reads{
    case _root_.scalapb.descriptors.PMessage(__fieldsMap) =>
      require(__fieldsMap.keys.forall(_.containingMessage == scalaDescriptor), "FieldDescriptor does not match message type.")
      payment.PaymentMethods(
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(2).get).map(_.as[_root_.scala.collection.Seq[payment.PaymentMethod]]).getOrElse(_root_.scala.collection.Seq.empty)
      )
    case _ => throw new RuntimeException("Expected PMessage")
  }
  def javaDescriptor: _root_.com.google.protobuf.Descriptors.Descriptor = PaymentProto.javaDescriptor.getMessageTypes.get(1)
  def scalaDescriptor: _root_.scalapb.descriptors.Descriptor = PaymentProto.scalaDescriptor.messages(1)
  def messageCompanionForFieldNumber(__number: _root_.scala.Int): _root_.scalapb.GeneratedMessageCompanion[_] = {
    var __out: _root_.scalapb.GeneratedMessageCompanion[_] = null
    (__number: @_root_.scala.unchecked) match {
      case 2 => __out = payment.PaymentMethod
    }
    __out
  }
  lazy val nestedMessagesCompanions: Seq[_root_.scalapb.GeneratedMessageCompanion[_]] = Seq.empty
  def enumCompanionForFieldNumber(__fieldNumber: _root_.scala.Int): _root_.scalapb.GeneratedEnumCompanion[_] = throw new MatchError(__fieldNumber)
  lazy val defaultInstance = payment.PaymentMethods(
  )
  implicit class PaymentMethodsLens[UpperPB](_l: _root_.scalapb.lenses.Lens[UpperPB, payment.PaymentMethods]) extends _root_.scalapb.lenses.ObjectLens[UpperPB, payment.PaymentMethods](_l) {
    def paymentMethods: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.collection.Seq[payment.PaymentMethod]] = field(_.paymentMethods)((c_, f_) => c_.copy(paymentMethods = f_))
  }
  final val PAYMENT_METHODS_FIELD_NUMBER = 2
}
