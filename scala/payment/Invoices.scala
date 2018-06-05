// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO2

package payment

@SerialVersionUID(0L)
final case class Invoices(
    invoices: _root_.scala.collection.Seq[payment.Invoice] = _root_.scala.collection.Seq.empty
    ) extends scalapb.GeneratedMessage with scalapb.Message[Invoices] with scalapb.lenses.Updatable[Invoices] {
    @transient
    private[this] var __serializedSizeCachedValue: _root_.scala.Int = 0
    private[this] def __computeSerializedValue(): _root_.scala.Int = {
      var __size = 0
      invoices.foreach(invoices => __size += 1 + _root_.com.google.protobuf.CodedOutputStream.computeUInt32SizeNoTag(invoices.serializedSize) + invoices.serializedSize)
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
      invoices.foreach { __v =>
        _output__.writeTag(2, 2)
        _output__.writeUInt32NoTag(__v.serializedSize)
        __v.writeTo(_output__)
      };
    }
    def mergeFrom(`_input__`: _root_.com.google.protobuf.CodedInputStream): payment.Invoices = {
      val __invoices = (_root_.scala.collection.immutable.Vector.newBuilder[payment.Invoice] ++= this.invoices)
      var _done__ = false
      while (!_done__) {
        val _tag__ = _input__.readTag()
        _tag__ match {
          case 0 => _done__ = true
          case 18 =>
            __invoices += _root_.scalapb.LiteParser.readMessage(_input__, payment.Invoice.defaultInstance)
          case tag => _input__.skipField(tag)
        }
      }
      payment.Invoices(
          invoices = __invoices.result()
      )
    }
    def clearInvoices = copy(invoices = _root_.scala.collection.Seq.empty)
    def addInvoices(__vs: payment.Invoice*): Invoices = addAllInvoices(__vs)
    def addAllInvoices(__vs: TraversableOnce[payment.Invoice]): Invoices = copy(invoices = invoices ++ __vs)
    def withInvoices(__v: _root_.scala.collection.Seq[payment.Invoice]): Invoices = copy(invoices = __v)
    def getFieldByNumber(__fieldNumber: _root_.scala.Int): scala.Any = {
      (__fieldNumber: @_root_.scala.unchecked) match {
        case 2 => invoices
      }
    }
    def getField(__field: _root_.scalapb.descriptors.FieldDescriptor): _root_.scalapb.descriptors.PValue = {
      require(__field.containingMessage eq companion.scalaDescriptor)
      (__field.number: @_root_.scala.unchecked) match {
        case 2 => _root_.scalapb.descriptors.PRepeated(invoices.map(_.toPMessage)(_root_.scala.collection.breakOut))
      }
    }
    def toProtoString: _root_.scala.Predef.String = _root_.scalapb.TextFormat.printToUnicodeString(this)
    def companion = payment.Invoices
}

object Invoices extends scalapb.GeneratedMessageCompanion[payment.Invoices] {
  implicit def messageCompanion: scalapb.GeneratedMessageCompanion[payment.Invoices] = this
  def fromFieldsMap(__fieldsMap: scala.collection.immutable.Map[_root_.com.google.protobuf.Descriptors.FieldDescriptor, scala.Any]): payment.Invoices = {
    require(__fieldsMap.keys.forall(_.getContainingType() == javaDescriptor), "FieldDescriptor does not match message type.")
    val __fields = javaDescriptor.getFields
    payment.Invoices(
      __fieldsMap.getOrElse(__fields.get(0), Nil).asInstanceOf[_root_.scala.collection.Seq[payment.Invoice]]
    )
  }
  implicit def messageReads: _root_.scalapb.descriptors.Reads[payment.Invoices] = _root_.scalapb.descriptors.Reads{
    case _root_.scalapb.descriptors.PMessage(__fieldsMap) =>
      require(__fieldsMap.keys.forall(_.containingMessage == scalaDescriptor), "FieldDescriptor does not match message type.")
      payment.Invoices(
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(2).get).map(_.as[_root_.scala.collection.Seq[payment.Invoice]]).getOrElse(_root_.scala.collection.Seq.empty)
      )
    case _ => throw new RuntimeException("Expected PMessage")
  }
  def javaDescriptor: _root_.com.google.protobuf.Descriptors.Descriptor = PaymentProto.javaDescriptor.getMessageTypes.get(9)
  def scalaDescriptor: _root_.scalapb.descriptors.Descriptor = PaymentProto.scalaDescriptor.messages(9)
  def messageCompanionForFieldNumber(__number: _root_.scala.Int): _root_.scalapb.GeneratedMessageCompanion[_] = {
    var __out: _root_.scalapb.GeneratedMessageCompanion[_] = null
    (__number: @_root_.scala.unchecked) match {
      case 2 => __out = payment.Invoice
    }
    __out
  }
  lazy val nestedMessagesCompanions: Seq[_root_.scalapb.GeneratedMessageCompanion[_]] = Seq.empty
  def enumCompanionForFieldNumber(__fieldNumber: _root_.scala.Int): _root_.scalapb.GeneratedEnumCompanion[_] = throw new MatchError(__fieldNumber)
  lazy val defaultInstance = payment.Invoices(
  )
  implicit class InvoicesLens[UpperPB](_l: _root_.scalapb.lenses.Lens[UpperPB, payment.Invoices]) extends _root_.scalapb.lenses.ObjectLens[UpperPB, payment.Invoices](_l) {
    def invoices: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.collection.Seq[payment.Invoice]] = field(_.invoices)((c_, f_) => c_.copy(invoices = f_))
  }
  final val INVOICES_FIELD_NUMBER = 2
}
