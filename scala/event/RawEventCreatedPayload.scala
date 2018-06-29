// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO3

package event

/** @param payload
  *  json format
  * @param payloads
  *   payload for each sub, payload must be empty
  */
@SerialVersionUID(0L)
final case class RawEventCreatedPayload(
    ctx: scala.Option[common.Context] = None,
    subs: _root_.scala.collection.Seq[_root_.scala.Predef.String] = _root_.scala.collection.Seq.empty,
    targetTopic: _root_.scala.Predef.String = "",
    payload: _root_.scala.Predef.String = "",
    targetKey: _root_.scala.Predef.String = "",
    payloads: _root_.scala.collection.Seq[_root_.scala.Predef.String] = _root_.scala.collection.Seq.empty,
    topic: _root_.scala.Predef.String = "",
    routerTopic: _root_.scala.Predef.String = "",
    sub: _root_.scala.Predef.String = ""
    ) extends scalapb.GeneratedMessage with scalapb.Message[RawEventCreatedPayload] with scalapb.lenses.Updatable[RawEventCreatedPayload] {
    @transient
    private[this] var __serializedSizeCachedValue: _root_.scala.Int = 0
    private[this] def __computeSerializedValue(): _root_.scala.Int = {
      var __size = 0
      if (ctx.isDefined) { __size += 1 + _root_.com.google.protobuf.CodedOutputStream.computeUInt32SizeNoTag(ctx.get.serializedSize) + ctx.get.serializedSize }
      subs.foreach(subs => __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(3, subs))
      if (targetTopic != "") { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(4, targetTopic) }
      if (payload != "") { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(5, payload) }
      if (targetKey != "") { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(6, targetKey) }
      payloads.foreach(payloads => __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(7, payloads))
      if (topic != "") { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(9, topic) }
      if (routerTopic != "") { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(10, routerTopic) }
      if (sub != "") { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(11, sub) }
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
      ctx.foreach { __v =>
        _output__.writeTag(1, 2)
        _output__.writeUInt32NoTag(__v.serializedSize)
        __v.writeTo(_output__)
      };
      subs.foreach { __v =>
        _output__.writeString(3, __v)
      };
      {
        val __v = targetTopic
        if (__v != "") {
          _output__.writeString(4, __v)
        }
      };
      {
        val __v = payload
        if (__v != "") {
          _output__.writeString(5, __v)
        }
      };
      {
        val __v = targetKey
        if (__v != "") {
          _output__.writeString(6, __v)
        }
      };
      payloads.foreach { __v =>
        _output__.writeString(7, __v)
      };
      {
        val __v = topic
        if (__v != "") {
          _output__.writeString(9, __v)
        }
      };
      {
        val __v = routerTopic
        if (__v != "") {
          _output__.writeString(10, __v)
        }
      };
      {
        val __v = sub
        if (__v != "") {
          _output__.writeString(11, __v)
        }
      };
    }
    def mergeFrom(`_input__`: _root_.com.google.protobuf.CodedInputStream): event.RawEventCreatedPayload = {
      var __ctx = this.ctx
      val __subs = (_root_.scala.collection.immutable.Vector.newBuilder[_root_.scala.Predef.String] ++= this.subs)
      var __targetTopic = this.targetTopic
      var __payload = this.payload
      var __targetKey = this.targetKey
      val __payloads = (_root_.scala.collection.immutable.Vector.newBuilder[_root_.scala.Predef.String] ++= this.payloads)
      var __topic = this.topic
      var __routerTopic = this.routerTopic
      var __sub = this.sub
      var _done__ = false
      while (!_done__) {
        val _tag__ = _input__.readTag()
        _tag__ match {
          case 0 => _done__ = true
          case 10 =>
            __ctx = Option(_root_.scalapb.LiteParser.readMessage(_input__, __ctx.getOrElse(common.Context.defaultInstance)))
          case 26 =>
            __subs += _input__.readString()
          case 34 =>
            __targetTopic = _input__.readString()
          case 42 =>
            __payload = _input__.readString()
          case 50 =>
            __targetKey = _input__.readString()
          case 58 =>
            __payloads += _input__.readString()
          case 74 =>
            __topic = _input__.readString()
          case 82 =>
            __routerTopic = _input__.readString()
          case 90 =>
            __sub = _input__.readString()
          case tag => _input__.skipField(tag)
        }
      }
      event.RawEventCreatedPayload(
          ctx = __ctx,
          subs = __subs.result(),
          targetTopic = __targetTopic,
          payload = __payload,
          targetKey = __targetKey,
          payloads = __payloads.result(),
          topic = __topic,
          routerTopic = __routerTopic,
          sub = __sub
      )
    }
    def getCtx: common.Context = ctx.getOrElse(common.Context.defaultInstance)
    def clearCtx: RawEventCreatedPayload = copy(ctx = None)
    def withCtx(__v: common.Context): RawEventCreatedPayload = copy(ctx = Option(__v))
    def clearSubs = copy(subs = _root_.scala.collection.Seq.empty)
    def addSubs(__vs: _root_.scala.Predef.String*): RawEventCreatedPayload = addAllSubs(__vs)
    def addAllSubs(__vs: TraversableOnce[_root_.scala.Predef.String]): RawEventCreatedPayload = copy(subs = subs ++ __vs)
    def withSubs(__v: _root_.scala.collection.Seq[_root_.scala.Predef.String]): RawEventCreatedPayload = copy(subs = __v)
    def withTargetTopic(__v: _root_.scala.Predef.String): RawEventCreatedPayload = copy(targetTopic = __v)
    def withPayload(__v: _root_.scala.Predef.String): RawEventCreatedPayload = copy(payload = __v)
    def withTargetKey(__v: _root_.scala.Predef.String): RawEventCreatedPayload = copy(targetKey = __v)
    def clearPayloads = copy(payloads = _root_.scala.collection.Seq.empty)
    def addPayloads(__vs: _root_.scala.Predef.String*): RawEventCreatedPayload = addAllPayloads(__vs)
    def addAllPayloads(__vs: TraversableOnce[_root_.scala.Predef.String]): RawEventCreatedPayload = copy(payloads = payloads ++ __vs)
    def withPayloads(__v: _root_.scala.collection.Seq[_root_.scala.Predef.String]): RawEventCreatedPayload = copy(payloads = __v)
    def withTopic(__v: _root_.scala.Predef.String): RawEventCreatedPayload = copy(topic = __v)
    def withRouterTopic(__v: _root_.scala.Predef.String): RawEventCreatedPayload = copy(routerTopic = __v)
    def withSub(__v: _root_.scala.Predef.String): RawEventCreatedPayload = copy(sub = __v)
    def getFieldByNumber(__fieldNumber: _root_.scala.Int): scala.Any = {
      (__fieldNumber: @_root_.scala.unchecked) match {
        case 1 => ctx.orNull
        case 3 => subs
        case 4 => {
          val __t = targetTopic
          if (__t != "") __t else null
        }
        case 5 => {
          val __t = payload
          if (__t != "") __t else null
        }
        case 6 => {
          val __t = targetKey
          if (__t != "") __t else null
        }
        case 7 => payloads
        case 9 => {
          val __t = topic
          if (__t != "") __t else null
        }
        case 10 => {
          val __t = routerTopic
          if (__t != "") __t else null
        }
        case 11 => {
          val __t = sub
          if (__t != "") __t else null
        }
      }
    }
    def getField(__field: _root_.scalapb.descriptors.FieldDescriptor): _root_.scalapb.descriptors.PValue = {
      require(__field.containingMessage eq companion.scalaDescriptor)
      (__field.number: @_root_.scala.unchecked) match {
        case 1 => ctx.map(_.toPMessage).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 3 => _root_.scalapb.descriptors.PRepeated(subs.map(_root_.scalapb.descriptors.PString)(_root_.scala.collection.breakOut))
        case 4 => _root_.scalapb.descriptors.PString(targetTopic)
        case 5 => _root_.scalapb.descriptors.PString(payload)
        case 6 => _root_.scalapb.descriptors.PString(targetKey)
        case 7 => _root_.scalapb.descriptors.PRepeated(payloads.map(_root_.scalapb.descriptors.PString)(_root_.scala.collection.breakOut))
        case 9 => _root_.scalapb.descriptors.PString(topic)
        case 10 => _root_.scalapb.descriptors.PString(routerTopic)
        case 11 => _root_.scalapb.descriptors.PString(sub)
      }
    }
    def toProtoString: _root_.scala.Predef.String = _root_.scalapb.TextFormat.printToUnicodeString(this)
    def companion = event.RawEventCreatedPayload
}

object RawEventCreatedPayload extends scalapb.GeneratedMessageCompanion[event.RawEventCreatedPayload] {
  implicit def messageCompanion: scalapb.GeneratedMessageCompanion[event.RawEventCreatedPayload] = this
  def fromFieldsMap(__fieldsMap: scala.collection.immutable.Map[_root_.com.google.protobuf.Descriptors.FieldDescriptor, scala.Any]): event.RawEventCreatedPayload = {
    require(__fieldsMap.keys.forall(_.getContainingType() == javaDescriptor), "FieldDescriptor does not match message type.")
    val __fields = javaDescriptor.getFields
    event.RawEventCreatedPayload(
      __fieldsMap.get(__fields.get(0)).asInstanceOf[scala.Option[common.Context]],
      __fieldsMap.getOrElse(__fields.get(1), Nil).asInstanceOf[_root_.scala.collection.Seq[_root_.scala.Predef.String]],
      __fieldsMap.getOrElse(__fields.get(2), "").asInstanceOf[_root_.scala.Predef.String],
      __fieldsMap.getOrElse(__fields.get(3), "").asInstanceOf[_root_.scala.Predef.String],
      __fieldsMap.getOrElse(__fields.get(4), "").asInstanceOf[_root_.scala.Predef.String],
      __fieldsMap.getOrElse(__fields.get(5), Nil).asInstanceOf[_root_.scala.collection.Seq[_root_.scala.Predef.String]],
      __fieldsMap.getOrElse(__fields.get(6), "").asInstanceOf[_root_.scala.Predef.String],
      __fieldsMap.getOrElse(__fields.get(7), "").asInstanceOf[_root_.scala.Predef.String],
      __fieldsMap.getOrElse(__fields.get(8), "").asInstanceOf[_root_.scala.Predef.String]
    )
  }
  implicit def messageReads: _root_.scalapb.descriptors.Reads[event.RawEventCreatedPayload] = _root_.scalapb.descriptors.Reads{
    case _root_.scalapb.descriptors.PMessage(__fieldsMap) =>
      require(__fieldsMap.keys.forall(_.containingMessage == scalaDescriptor), "FieldDescriptor does not match message type.")
      event.RawEventCreatedPayload(
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(1).get).flatMap(_.as[scala.Option[common.Context]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(3).get).map(_.as[_root_.scala.collection.Seq[_root_.scala.Predef.String]]).getOrElse(_root_.scala.collection.Seq.empty),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(4).get).map(_.as[_root_.scala.Predef.String]).getOrElse(""),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(5).get).map(_.as[_root_.scala.Predef.String]).getOrElse(""),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(6).get).map(_.as[_root_.scala.Predef.String]).getOrElse(""),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(7).get).map(_.as[_root_.scala.collection.Seq[_root_.scala.Predef.String]]).getOrElse(_root_.scala.collection.Seq.empty),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(9).get).map(_.as[_root_.scala.Predef.String]).getOrElse(""),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(10).get).map(_.as[_root_.scala.Predef.String]).getOrElse(""),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(11).get).map(_.as[_root_.scala.Predef.String]).getOrElse("")
      )
    case _ => throw new RuntimeException("Expected PMessage")
  }
  def javaDescriptor: _root_.com.google.protobuf.Descriptors.Descriptor = EventProto.javaDescriptor.getMessageTypes.get(0)
  def scalaDescriptor: _root_.scalapb.descriptors.Descriptor = EventProto.scalaDescriptor.messages(0)
  def messageCompanionForFieldNumber(__number: _root_.scala.Int): _root_.scalapb.GeneratedMessageCompanion[_] = {
    var __out: _root_.scalapb.GeneratedMessageCompanion[_] = null
    (__number: @_root_.scala.unchecked) match {
      case 1 => __out = common.Context
    }
    __out
  }
  lazy val nestedMessagesCompanions: Seq[_root_.scalapb.GeneratedMessageCompanion[_]] = Seq.empty
  def enumCompanionForFieldNumber(__fieldNumber: _root_.scala.Int): _root_.scalapb.GeneratedEnumCompanion[_] = throw new MatchError(__fieldNumber)
  lazy val defaultInstance = event.RawEventCreatedPayload(
  )
  implicit class RawEventCreatedPayloadLens[UpperPB](_l: _root_.scalapb.lenses.Lens[UpperPB, event.RawEventCreatedPayload]) extends _root_.scalapb.lenses.ObjectLens[UpperPB, event.RawEventCreatedPayload](_l) {
    def ctx: _root_.scalapb.lenses.Lens[UpperPB, common.Context] = field(_.getCtx)((c_, f_) => c_.copy(ctx = Option(f_)))
    def optionalCtx: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[common.Context]] = field(_.ctx)((c_, f_) => c_.copy(ctx = f_))
    def subs: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.collection.Seq[_root_.scala.Predef.String]] = field(_.subs)((c_, f_) => c_.copy(subs = f_))
    def targetTopic: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.targetTopic)((c_, f_) => c_.copy(targetTopic = f_))
    def payload: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.payload)((c_, f_) => c_.copy(payload = f_))
    def targetKey: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.targetKey)((c_, f_) => c_.copy(targetKey = f_))
    def payloads: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.collection.Seq[_root_.scala.Predef.String]] = field(_.payloads)((c_, f_) => c_.copy(payloads = f_))
    def topic: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.topic)((c_, f_) => c_.copy(topic = f_))
    def routerTopic: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.routerTopic)((c_, f_) => c_.copy(routerTopic = f_))
    def sub: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.sub)((c_, f_) => c_.copy(sub = f_))
  }
  final val CTX_FIELD_NUMBER = 1
  final val SUBS_FIELD_NUMBER = 3
  final val TARGET_TOPIC_FIELD_NUMBER = 4
  final val PAYLOAD_FIELD_NUMBER = 5
  final val TARGET_KEY_FIELD_NUMBER = 6
  final val PAYLOADS_FIELD_NUMBER = 7
  final val TOPIC_FIELD_NUMBER = 9
  final val ROUTER_TOPIC_FIELD_NUMBER = 10
  final val SUB_FIELD_NUMBER = 11
}