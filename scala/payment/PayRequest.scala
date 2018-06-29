// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO2

package payment

@SerialVersionUID(0L)
final case class PayRequest(
    accountId: scala.Option[_root_.scala.Predef.String] = None,
    invoiceIds: _root_.scala.collection.Seq[_root_.scala.Predef.String] = _root_.scala.collection.Seq.empty,
    description: scala.Option[_root_.scala.Predef.String] = None,
    customerInfo: scala.Option[payment.Contact] = None,
    amount: scala.Option[_root_.scala.Float] = None
    ) extends scalapb.GeneratedMessage with scalapb.Message[PayRequest] with scalapb.lenses.Updatable[PayRequest] {
    @transient
    private[this] var __serializedSizeCachedValue: _root_.scala.Int = 0
    private[this] def __computeSerializedValue(): _root_.scala.Int = {
      var __size = 0
      if (accountId.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(8, accountId.get) }
      invoiceIds.foreach(invoiceIds => __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(6, invoiceIds))
      if (description.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(7, description.get) }
      if (customerInfo.isDefined) { __size += 1 + _root_.com.google.protobuf.CodedOutputStream.computeUInt32SizeNoTag(customerInfo.get.serializedSize) + customerInfo.get.serializedSize }
      if (amount.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeFloatSize(10, amount.get) }
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
      invoiceIds.foreach { __v =>
        _output__.writeString(6, __v)
      };
      description.foreach { __v =>
        _output__.writeString(7, __v)
      };
      accountId.foreach { __v =>
        _output__.writeString(8, __v)
      };
      customerInfo.foreach { __v =>
        _output__.writeTag(9, 2)
        _output__.writeUInt32NoTag(__v.serializedSize)
        __v.writeTo(_output__)
      };
      amount.foreach { __v =>
        _output__.writeFloat(10, __v)
      };
    }
    def mergeFrom(`_input__`: _root_.com.google.protobuf.CodedInputStream): payment.PayRequest = {
      var __accountId = this.accountId
      val __invoiceIds = (_root_.scala.collection.immutable.Vector.newBuilder[_root_.scala.Predef.String] ++= this.invoiceIds)
      var __description = this.description
      var __customerInfo = this.customerInfo
      var __amount = this.amount
      var _done__ = false
      while (!_done__) {
        val _tag__ = _input__.readTag()
        _tag__ match {
          case 0 => _done__ = true
          case 66 =>
            __accountId = Option(_input__.readString())
          case 50 =>
            __invoiceIds += _input__.readString()
          case 58 =>
            __description = Option(_input__.readString())
          case 74 =>
            __customerInfo = Option(_root_.scalapb.LiteParser.readMessage(_input__, __customerInfo.getOrElse(payment.Contact.defaultInstance)))
          case 85 =>
            __amount = Option(_input__.readFloat())
          case tag => _input__.skipField(tag)
        }
      }
      payment.PayRequest(
          accountId = __accountId,
          invoiceIds = __invoiceIds.result(),
          description = __description,
          customerInfo = __customerInfo,
          amount = __amount
      )
    }
    def getAccountId: _root_.scala.Predef.String = accountId.getOrElse("")
    def clearAccountId: PayRequest = copy(accountId = None)
    def withAccountId(__v: _root_.scala.Predef.String): PayRequest = copy(accountId = Option(__v))
    def clearInvoiceIds = copy(invoiceIds = _root_.scala.collection.Seq.empty)
    def addInvoiceIds(__vs: _root_.scala.Predef.String*): PayRequest = addAllInvoiceIds(__vs)
    def addAllInvoiceIds(__vs: TraversableOnce[_root_.scala.Predef.String]): PayRequest = copy(invoiceIds = invoiceIds ++ __vs)
    def withInvoiceIds(__v: _root_.scala.collection.Seq[_root_.scala.Predef.String]): PayRequest = copy(invoiceIds = __v)
    def getDescription: _root_.scala.Predef.String = description.getOrElse("")
    def clearDescription: PayRequest = copy(description = None)
    def withDescription(__v: _root_.scala.Predef.String): PayRequest = copy(description = Option(__v))
    def getCustomerInfo: payment.Contact = customerInfo.getOrElse(payment.Contact.defaultInstance)
    def clearCustomerInfo: PayRequest = copy(customerInfo = None)
    def withCustomerInfo(__v: payment.Contact): PayRequest = copy(customerInfo = Option(__v))
    def getAmount: _root_.scala.Float = amount.getOrElse(0.0f)
    def clearAmount: PayRequest = copy(amount = None)
    def withAmount(__v: _root_.scala.Float): PayRequest = copy(amount = Option(__v))
    def getFieldByNumber(__fieldNumber: _root_.scala.Int): scala.Any = {
      (__fieldNumber: @_root_.scala.unchecked) match {
        case 8 => accountId.orNull
        case 6 => invoiceIds
        case 7 => description.orNull
        case 9 => customerInfo.orNull
        case 10 => amount.orNull
      }
    }
    def getField(__field: _root_.scalapb.descriptors.FieldDescriptor): _root_.scalapb.descriptors.PValue = {
      require(__field.containingMessage eq companion.scalaDescriptor)
      (__field.number: @_root_.scala.unchecked) match {
        case 8 => accountId.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 6 => _root_.scalapb.descriptors.PRepeated(invoiceIds.map(_root_.scalapb.descriptors.PString)(_root_.scala.collection.breakOut))
        case 7 => description.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 9 => customerInfo.map(_.toPMessage).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 10 => amount.map(_root_.scalapb.descriptors.PFloat).getOrElse(_root_.scalapb.descriptors.PEmpty)
      }
    }
    def toProtoString: _root_.scala.Predef.String = _root_.scalapb.TextFormat.printToUnicodeString(this)
    def companion = payment.PayRequest
}

object PayRequest extends scalapb.GeneratedMessageCompanion[payment.PayRequest] {
  implicit def messageCompanion: scalapb.GeneratedMessageCompanion[payment.PayRequest] = this
  def fromFieldsMap(__fieldsMap: scala.collection.immutable.Map[_root_.com.google.protobuf.Descriptors.FieldDescriptor, scala.Any]): payment.PayRequest = {
    require(__fieldsMap.keys.forall(_.getContainingType() == javaDescriptor), "FieldDescriptor does not match message type.")
    val __fields = javaDescriptor.getFields
    payment.PayRequest(
      __fieldsMap.get(__fields.get(0)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.getOrElse(__fields.get(1), Nil).asInstanceOf[_root_.scala.collection.Seq[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(2)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(3)).asInstanceOf[scala.Option[payment.Contact]],
      __fieldsMap.get(__fields.get(4)).asInstanceOf[scala.Option[_root_.scala.Float]]
    )
  }
  implicit def messageReads: _root_.scalapb.descriptors.Reads[payment.PayRequest] = _root_.scalapb.descriptors.Reads{
    case _root_.scalapb.descriptors.PMessage(__fieldsMap) =>
      require(__fieldsMap.keys.forall(_.containingMessage == scalaDescriptor), "FieldDescriptor does not match message type.")
      payment.PayRequest(
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(8).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(6).get).map(_.as[_root_.scala.collection.Seq[_root_.scala.Predef.String]]).getOrElse(_root_.scala.collection.Seq.empty),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(7).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(9).get).flatMap(_.as[scala.Option[payment.Contact]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(10).get).flatMap(_.as[scala.Option[_root_.scala.Float]])
      )
    case _ => throw new RuntimeException("Expected PMessage")
  }
  def javaDescriptor: _root_.com.google.protobuf.Descriptors.Descriptor = PaymentProto.javaDescriptor.getMessageTypes.get(28)
  def scalaDescriptor: _root_.scalapb.descriptors.Descriptor = PaymentProto.scalaDescriptor.messages(28)
  def messageCompanionForFieldNumber(__number: _root_.scala.Int): _root_.scalapb.GeneratedMessageCompanion[_] = {
    var __out: _root_.scalapb.GeneratedMessageCompanion[_] = null
    (__number: @_root_.scala.unchecked) match {
      case 9 => __out = payment.Contact
    }
    __out
  }
  lazy val nestedMessagesCompanions: Seq[_root_.scalapb.GeneratedMessageCompanion[_]] = Seq.empty
  def enumCompanionForFieldNumber(__fieldNumber: _root_.scala.Int): _root_.scalapb.GeneratedEnumCompanion[_] = throw new MatchError(__fieldNumber)
  lazy val defaultInstance = payment.PayRequest(
  )
  implicit class PayRequestLens[UpperPB](_l: _root_.scalapb.lenses.Lens[UpperPB, payment.PayRequest]) extends _root_.scalapb.lenses.ObjectLens[UpperPB, payment.PayRequest](_l) {
    def accountId: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getAccountId)((c_, f_) => c_.copy(accountId = Option(f_)))
    def optionalAccountId: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.accountId)((c_, f_) => c_.copy(accountId = f_))
    def invoiceIds: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.collection.Seq[_root_.scala.Predef.String]] = field(_.invoiceIds)((c_, f_) => c_.copy(invoiceIds = f_))
    def description: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getDescription)((c_, f_) => c_.copy(description = Option(f_)))
    def optionalDescription: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.description)((c_, f_) => c_.copy(description = f_))
    def customerInfo: _root_.scalapb.lenses.Lens[UpperPB, payment.Contact] = field(_.getCustomerInfo)((c_, f_) => c_.copy(customerInfo = Option(f_)))
    def optionalCustomerInfo: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[payment.Contact]] = field(_.customerInfo)((c_, f_) => c_.copy(customerInfo = f_))
    def amount: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Float] = field(_.getAmount)((c_, f_) => c_.copy(amount = Option(f_)))
    def optionalAmount: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Float]] = field(_.amount)((c_, f_) => c_.copy(amount = f_))
  }
  final val ACCOUNT_ID_FIELD_NUMBER = 8
  final val INVOICE_IDS_FIELD_NUMBER = 6
  final val DESCRIPTION_FIELD_NUMBER = 7
  final val CUSTOMERINFO_FIELD_NUMBER = 9
  final val AMOUNT_FIELD_NUMBER = 10
}