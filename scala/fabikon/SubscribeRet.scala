// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO2

package fabikon

@SerialVersionUID(0L)
final case class SubscribeRet(
    success: scala.Option[_root_.scala.Boolean] = None
    ) extends scalapb.GeneratedMessage with scalapb.Message[SubscribeRet] with scalapb.lenses.Updatable[SubscribeRet] {
    @transient
    private[this] var __serializedSizeCachedValue: _root_.scala.Int = 0
    private[this] def __computeSerializedValue(): _root_.scala.Int = {
      var __size = 0
      if (success.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeBoolSize(2, success.get) }
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
      success.foreach { __v =>
        _output__.writeBool(2, __v)
      };
    }
    def mergeFrom(`_input__`: _root_.com.google.protobuf.CodedInputStream): fabikon.SubscribeRet = {
      var __success = this.success
      var _done__ = false
      while (!_done__) {
        val _tag__ = _input__.readTag()
        _tag__ match {
          case 0 => _done__ = true
          case 16 =>
            __success = Option(_input__.readBool())
          case tag => _input__.skipField(tag)
        }
      }
      fabikon.SubscribeRet(
          success = __success
      )
    }
    def getSuccess: _root_.scala.Boolean = success.getOrElse(false)
    def clearSuccess: SubscribeRet = copy(success = None)
    def withSuccess(__v: _root_.scala.Boolean): SubscribeRet = copy(success = Option(__v))
    def getFieldByNumber(__fieldNumber: _root_.scala.Int): scala.Any = {
      (__fieldNumber: @_root_.scala.unchecked) match {
        case 2 => success.orNull
      }
    }
    def getField(__field: _root_.scalapb.descriptors.FieldDescriptor): _root_.scalapb.descriptors.PValue = {
      require(__field.containingMessage eq companion.scalaDescriptor)
      (__field.number: @_root_.scala.unchecked) match {
        case 2 => success.map(_root_.scalapb.descriptors.PBoolean).getOrElse(_root_.scalapb.descriptors.PEmpty)
      }
    }
    def toProtoString: _root_.scala.Predef.String = _root_.scalapb.TextFormat.printToUnicodeString(this)
    def companion = fabikon.SubscribeRet
}

object SubscribeRet extends scalapb.GeneratedMessageCompanion[fabikon.SubscribeRet] {
  implicit def messageCompanion: scalapb.GeneratedMessageCompanion[fabikon.SubscribeRet] = this
  def fromFieldsMap(__fieldsMap: scala.collection.immutable.Map[_root_.com.google.protobuf.Descriptors.FieldDescriptor, scala.Any]): fabikon.SubscribeRet = {
    require(__fieldsMap.keys.forall(_.getContainingType() == javaDescriptor), "FieldDescriptor does not match message type.")
    val __fields = javaDescriptor.getFields
    fabikon.SubscribeRet(
      __fieldsMap.get(__fields.get(0)).asInstanceOf[scala.Option[_root_.scala.Boolean]]
    )
  }
  implicit def messageReads: _root_.scalapb.descriptors.Reads[fabikon.SubscribeRet] = _root_.scalapb.descriptors.Reads{
    case _root_.scalapb.descriptors.PMessage(__fieldsMap) =>
      require(__fieldsMap.keys.forall(_.containingMessage == scalaDescriptor), "FieldDescriptor does not match message type.")
      fabikon.SubscribeRet(
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(2).get).flatMap(_.as[scala.Option[_root_.scala.Boolean]])
      )
    case _ => throw new RuntimeException("Expected PMessage")
  }
  def javaDescriptor: _root_.com.google.protobuf.Descriptors.Descriptor = FabikonProto.javaDescriptor.getMessageTypes.get(24)
  def scalaDescriptor: _root_.scalapb.descriptors.Descriptor = FabikonProto.scalaDescriptor.messages(24)
  def messageCompanionForFieldNumber(__number: _root_.scala.Int): _root_.scalapb.GeneratedMessageCompanion[_] = throw new MatchError(__number)
  lazy val nestedMessagesCompanions: Seq[_root_.scalapb.GeneratedMessageCompanion[_]] = Seq.empty
  def enumCompanionForFieldNumber(__fieldNumber: _root_.scala.Int): _root_.scalapb.GeneratedEnumCompanion[_] = throw new MatchError(__fieldNumber)
  lazy val defaultInstance = fabikon.SubscribeRet(
  )
  implicit class SubscribeRetLens[UpperPB](_l: _root_.scalapb.lenses.Lens[UpperPB, fabikon.SubscribeRet]) extends _root_.scalapb.lenses.ObjectLens[UpperPB, fabikon.SubscribeRet](_l) {
    def success: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Boolean] = field(_.getSuccess)((c_, f_) => c_.copy(success = Option(f_)))
    def optionalSuccess: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Boolean]] = field(_.success)((c_, f_) => c_.copy(success = f_))
  }
  final val SUCCESS_FIELD_NUMBER = 2
}
