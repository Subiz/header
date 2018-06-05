// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO3

package ws

/** @param events
  *  common.Context ctx = 1;
  *  optional string connection_id = 2;
  */
@SerialVersionUID(0L)
final case class Subscribe(
    events: _root_.scala.collection.Seq[_root_.scala.Predef.String] = _root_.scala.collection.Seq.empty,
    connectionId: _root_.scala.Predef.String = ""
    ) extends scalapb.GeneratedMessage with scalapb.Message[Subscribe] with scalapb.lenses.Updatable[Subscribe] {
    @transient
    private[this] var __serializedSizeCachedValue: _root_.scala.Int = 0
    private[this] def __computeSerializedValue(): _root_.scala.Int = {
      var __size = 0
      events.foreach(events => __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(3, events))
      if (connectionId != "") { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(5, connectionId) }
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
      events.foreach { __v =>
        _output__.writeString(3, __v)
      };
      {
        val __v = connectionId
        if (__v != "") {
          _output__.writeString(5, __v)
        }
      };
    }
    def mergeFrom(`_input__`: _root_.com.google.protobuf.CodedInputStream): ws.Subscribe = {
      val __events = (_root_.scala.collection.immutable.Vector.newBuilder[_root_.scala.Predef.String] ++= this.events)
      var __connectionId = this.connectionId
      var _done__ = false
      while (!_done__) {
        val _tag__ = _input__.readTag()
        _tag__ match {
          case 0 => _done__ = true
          case 26 =>
            __events += _input__.readString()
          case 42 =>
            __connectionId = _input__.readString()
          case tag => _input__.skipField(tag)
        }
      }
      ws.Subscribe(
          events = __events.result(),
          connectionId = __connectionId
      )
    }
    def clearEvents = copy(events = _root_.scala.collection.Seq.empty)
    def addEvents(__vs: _root_.scala.Predef.String*): Subscribe = addAllEvents(__vs)
    def addAllEvents(__vs: TraversableOnce[_root_.scala.Predef.String]): Subscribe = copy(events = events ++ __vs)
    def withEvents(__v: _root_.scala.collection.Seq[_root_.scala.Predef.String]): Subscribe = copy(events = __v)
    def withConnectionId(__v: _root_.scala.Predef.String): Subscribe = copy(connectionId = __v)
    def getFieldByNumber(__fieldNumber: _root_.scala.Int): scala.Any = {
      (__fieldNumber: @_root_.scala.unchecked) match {
        case 3 => events
        case 5 => {
          val __t = connectionId
          if (__t != "") __t else null
        }
      }
    }
    def getField(__field: _root_.scalapb.descriptors.FieldDescriptor): _root_.scalapb.descriptors.PValue = {
      require(__field.containingMessage eq companion.scalaDescriptor)
      (__field.number: @_root_.scala.unchecked) match {
        case 3 => _root_.scalapb.descriptors.PRepeated(events.map(_root_.scalapb.descriptors.PString)(_root_.scala.collection.breakOut))
        case 5 => _root_.scalapb.descriptors.PString(connectionId)
      }
    }
    def toProtoString: _root_.scala.Predef.String = _root_.scalapb.TextFormat.printToUnicodeString(this)
    def companion = ws.Subscribe
}

object Subscribe extends scalapb.GeneratedMessageCompanion[ws.Subscribe] {
  implicit def messageCompanion: scalapb.GeneratedMessageCompanion[ws.Subscribe] = this
  def fromFieldsMap(__fieldsMap: scala.collection.immutable.Map[_root_.com.google.protobuf.Descriptors.FieldDescriptor, scala.Any]): ws.Subscribe = {
    require(__fieldsMap.keys.forall(_.getContainingType() == javaDescriptor), "FieldDescriptor does not match message type.")
    val __fields = javaDescriptor.getFields
    ws.Subscribe(
      __fieldsMap.getOrElse(__fields.get(0), Nil).asInstanceOf[_root_.scala.collection.Seq[_root_.scala.Predef.String]],
      __fieldsMap.getOrElse(__fields.get(1), "").asInstanceOf[_root_.scala.Predef.String]
    )
  }
  implicit def messageReads: _root_.scalapb.descriptors.Reads[ws.Subscribe] = _root_.scalapb.descriptors.Reads{
    case _root_.scalapb.descriptors.PMessage(__fieldsMap) =>
      require(__fieldsMap.keys.forall(_.containingMessage == scalaDescriptor), "FieldDescriptor does not match message type.")
      ws.Subscribe(
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(3).get).map(_.as[_root_.scala.collection.Seq[_root_.scala.Predef.String]]).getOrElse(_root_.scala.collection.Seq.empty),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(5).get).map(_.as[_root_.scala.Predef.String]).getOrElse("")
      )
    case _ => throw new RuntimeException("Expected PMessage")
  }
  def javaDescriptor: _root_.com.google.protobuf.Descriptors.Descriptor = WsProto.javaDescriptor.getMessageTypes.get(1)
  def scalaDescriptor: _root_.scalapb.descriptors.Descriptor = WsProto.scalaDescriptor.messages(1)
  def messageCompanionForFieldNumber(__number: _root_.scala.Int): _root_.scalapb.GeneratedMessageCompanion[_] = throw new MatchError(__number)
  lazy val nestedMessagesCompanions: Seq[_root_.scalapb.GeneratedMessageCompanion[_]] = Seq.empty
  def enumCompanionForFieldNumber(__fieldNumber: _root_.scala.Int): _root_.scalapb.GeneratedEnumCompanion[_] = throw new MatchError(__fieldNumber)
  lazy val defaultInstance = ws.Subscribe(
  )
  implicit class SubscribeLens[UpperPB](_l: _root_.scalapb.lenses.Lens[UpperPB, ws.Subscribe]) extends _root_.scalapb.lenses.ObjectLens[UpperPB, ws.Subscribe](_l) {
    def events: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.collection.Seq[_root_.scala.Predef.String]] = field(_.events)((c_, f_) => c_.copy(events = f_))
    def connectionId: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.connectionId)((c_, f_) => c_.copy(connectionId = f_))
  }
  final val EVENTS_FIELD_NUMBER = 3
  final val CONNECTION_ID_FIELD_NUMBER = 5
}
