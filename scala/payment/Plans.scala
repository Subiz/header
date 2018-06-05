// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO2

package payment

@SerialVersionUID(0L)
final case class Plans(
    plans: _root_.scala.collection.Seq[payment.Plan] = _root_.scala.collection.Seq.empty
    ) extends scalapb.GeneratedMessage with scalapb.Message[Plans] with scalapb.lenses.Updatable[Plans] {
    @transient
    private[this] var __serializedSizeCachedValue: _root_.scala.Int = 0
    private[this] def __computeSerializedValue(): _root_.scala.Int = {
      var __size = 0
      plans.foreach(plans => __size += 1 + _root_.com.google.protobuf.CodedOutputStream.computeUInt32SizeNoTag(plans.serializedSize) + plans.serializedSize)
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
      plans.foreach { __v =>
        _output__.writeTag(2, 2)
        _output__.writeUInt32NoTag(__v.serializedSize)
        __v.writeTo(_output__)
      };
    }
    def mergeFrom(`_input__`: _root_.com.google.protobuf.CodedInputStream): payment.Plans = {
      val __plans = (_root_.scala.collection.immutable.Vector.newBuilder[payment.Plan] ++= this.plans)
      var _done__ = false
      while (!_done__) {
        val _tag__ = _input__.readTag()
        _tag__ match {
          case 0 => _done__ = true
          case 18 =>
            __plans += _root_.scalapb.LiteParser.readMessage(_input__, payment.Plan.defaultInstance)
          case tag => _input__.skipField(tag)
        }
      }
      payment.Plans(
          plans = __plans.result()
      )
    }
    def clearPlans = copy(plans = _root_.scala.collection.Seq.empty)
    def addPlans(__vs: payment.Plan*): Plans = addAllPlans(__vs)
    def addAllPlans(__vs: TraversableOnce[payment.Plan]): Plans = copy(plans = plans ++ __vs)
    def withPlans(__v: _root_.scala.collection.Seq[payment.Plan]): Plans = copy(plans = __v)
    def getFieldByNumber(__fieldNumber: _root_.scala.Int): scala.Any = {
      (__fieldNumber: @_root_.scala.unchecked) match {
        case 2 => plans
      }
    }
    def getField(__field: _root_.scalapb.descriptors.FieldDescriptor): _root_.scalapb.descriptors.PValue = {
      require(__field.containingMessage eq companion.scalaDescriptor)
      (__field.number: @_root_.scala.unchecked) match {
        case 2 => _root_.scalapb.descriptors.PRepeated(plans.map(_.toPMessage)(_root_.scala.collection.breakOut))
      }
    }
    def toProtoString: _root_.scala.Predef.String = _root_.scalapb.TextFormat.printToUnicodeString(this)
    def companion = payment.Plans
}

object Plans extends scalapb.GeneratedMessageCompanion[payment.Plans] {
  implicit def messageCompanion: scalapb.GeneratedMessageCompanion[payment.Plans] = this
  def fromFieldsMap(__fieldsMap: scala.collection.immutable.Map[_root_.com.google.protobuf.Descriptors.FieldDescriptor, scala.Any]): payment.Plans = {
    require(__fieldsMap.keys.forall(_.getContainingType() == javaDescriptor), "FieldDescriptor does not match message type.")
    val __fields = javaDescriptor.getFields
    payment.Plans(
      __fieldsMap.getOrElse(__fields.get(0), Nil).asInstanceOf[_root_.scala.collection.Seq[payment.Plan]]
    )
  }
  implicit def messageReads: _root_.scalapb.descriptors.Reads[payment.Plans] = _root_.scalapb.descriptors.Reads{
    case _root_.scalapb.descriptors.PMessage(__fieldsMap) =>
      require(__fieldsMap.keys.forall(_.containingMessage == scalaDescriptor), "FieldDescriptor does not match message type.")
      payment.Plans(
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(2).get).map(_.as[_root_.scala.collection.Seq[payment.Plan]]).getOrElse(_root_.scala.collection.Seq.empty)
      )
    case _ => throw new RuntimeException("Expected PMessage")
  }
  def javaDescriptor: _root_.com.google.protobuf.Descriptors.Descriptor = PaymentProto.javaDescriptor.getMessageTypes.get(4)
  def scalaDescriptor: _root_.scalapb.descriptors.Descriptor = PaymentProto.scalaDescriptor.messages(4)
  def messageCompanionForFieldNumber(__number: _root_.scala.Int): _root_.scalapb.GeneratedMessageCompanion[_] = {
    var __out: _root_.scalapb.GeneratedMessageCompanion[_] = null
    (__number: @_root_.scala.unchecked) match {
      case 2 => __out = payment.Plan
    }
    __out
  }
  lazy val nestedMessagesCompanions: Seq[_root_.scalapb.GeneratedMessageCompanion[_]] = Seq.empty
  def enumCompanionForFieldNumber(__fieldNumber: _root_.scala.Int): _root_.scalapb.GeneratedEnumCompanion[_] = throw new MatchError(__fieldNumber)
  lazy val defaultInstance = payment.Plans(
  )
  implicit class PlansLens[UpperPB](_l: _root_.scalapb.lenses.Lens[UpperPB, payment.Plans]) extends _root_.scalapb.lenses.ObjectLens[UpperPB, payment.Plans](_l) {
    def plans: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.collection.Seq[payment.Plan]] = field(_.plans)((c_, f_) => c_.copy(plans = f_))
  }
  final val PLANS_FIELD_NUMBER = 2
}
