// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO2

package user

/** @param agentId
  *   search my user of agent
  * @param unread
  *   search my user of agent
  */
@SerialVersionUID(0L)
final case class UserSearchRequest(
    ctx: scala.Option[common.Context] = None,
    segmentationId: scala.Option[_root_.scala.Predef.String] = None,
    query: scala.Option[_root_.scala.Predef.String] = None,
    anchor: scala.Option[_root_.scala.Predef.String] = None,
    limit: scala.Option[_root_.scala.Int] = None,
    agentId: scala.Option[_root_.scala.Predef.String] = None,
    unread: scala.Option[_root_.scala.Boolean] = None
    ) extends scalapb.GeneratedMessage with scalapb.Message[UserSearchRequest] with scalapb.lenses.Updatable[UserSearchRequest] {
    @transient
    private[this] var __serializedSizeCachedValue: _root_.scala.Int = 0
    private[this] def __computeSerializedValue(): _root_.scala.Int = {
      var __size = 0
      if (ctx.isDefined) { __size += 1 + _root_.com.google.protobuf.CodedOutputStream.computeUInt32SizeNoTag(ctx.get.serializedSize) + ctx.get.serializedSize }
      if (segmentationId.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(3, segmentationId.get) }
      if (query.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(4, query.get) }
      if (anchor.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(5, anchor.get) }
      if (limit.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeInt32Size(6, limit.get) }
      if (agentId.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(8, agentId.get) }
      if (unread.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeBoolSize(9, unread.get) }
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
      segmentationId.foreach { __v =>
        _output__.writeString(3, __v)
      };
      query.foreach { __v =>
        _output__.writeString(4, __v)
      };
      anchor.foreach { __v =>
        _output__.writeString(5, __v)
      };
      limit.foreach { __v =>
        _output__.writeInt32(6, __v)
      };
      agentId.foreach { __v =>
        _output__.writeString(8, __v)
      };
      unread.foreach { __v =>
        _output__.writeBool(9, __v)
      };
    }
    def mergeFrom(`_input__`: _root_.com.google.protobuf.CodedInputStream): user.UserSearchRequest = {
      var __ctx = this.ctx
      var __segmentationId = this.segmentationId
      var __query = this.query
      var __anchor = this.anchor
      var __limit = this.limit
      var __agentId = this.agentId
      var __unread = this.unread
      var _done__ = false
      while (!_done__) {
        val _tag__ = _input__.readTag()
        _tag__ match {
          case 0 => _done__ = true
          case 10 =>
            __ctx = Option(_root_.scalapb.LiteParser.readMessage(_input__, __ctx.getOrElse(common.Context.defaultInstance)))
          case 26 =>
            __segmentationId = Option(_input__.readString())
          case 34 =>
            __query = Option(_input__.readString())
          case 42 =>
            __anchor = Option(_input__.readString())
          case 48 =>
            __limit = Option(_input__.readInt32())
          case 66 =>
            __agentId = Option(_input__.readString())
          case 72 =>
            __unread = Option(_input__.readBool())
          case tag => _input__.skipField(tag)
        }
      }
      user.UserSearchRequest(
          ctx = __ctx,
          segmentationId = __segmentationId,
          query = __query,
          anchor = __anchor,
          limit = __limit,
          agentId = __agentId,
          unread = __unread
      )
    }
    def getCtx: common.Context = ctx.getOrElse(common.Context.defaultInstance)
    def clearCtx: UserSearchRequest = copy(ctx = None)
    def withCtx(__v: common.Context): UserSearchRequest = copy(ctx = Option(__v))
    def getSegmentationId: _root_.scala.Predef.String = segmentationId.getOrElse("")
    def clearSegmentationId: UserSearchRequest = copy(segmentationId = None)
    def withSegmentationId(__v: _root_.scala.Predef.String): UserSearchRequest = copy(segmentationId = Option(__v))
    def getQuery: _root_.scala.Predef.String = query.getOrElse("")
    def clearQuery: UserSearchRequest = copy(query = None)
    def withQuery(__v: _root_.scala.Predef.String): UserSearchRequest = copy(query = Option(__v))
    def getAnchor: _root_.scala.Predef.String = anchor.getOrElse("")
    def clearAnchor: UserSearchRequest = copy(anchor = None)
    def withAnchor(__v: _root_.scala.Predef.String): UserSearchRequest = copy(anchor = Option(__v))
    def getLimit: _root_.scala.Int = limit.getOrElse(0)
    def clearLimit: UserSearchRequest = copy(limit = None)
    def withLimit(__v: _root_.scala.Int): UserSearchRequest = copy(limit = Option(__v))
    def getAgentId: _root_.scala.Predef.String = agentId.getOrElse("")
    def clearAgentId: UserSearchRequest = copy(agentId = None)
    def withAgentId(__v: _root_.scala.Predef.String): UserSearchRequest = copy(agentId = Option(__v))
    def getUnread: _root_.scala.Boolean = unread.getOrElse(false)
    def clearUnread: UserSearchRequest = copy(unread = None)
    def withUnread(__v: _root_.scala.Boolean): UserSearchRequest = copy(unread = Option(__v))
    def getFieldByNumber(__fieldNumber: _root_.scala.Int): scala.Any = {
      (__fieldNumber: @_root_.scala.unchecked) match {
        case 1 => ctx.orNull
        case 3 => segmentationId.orNull
        case 4 => query.orNull
        case 5 => anchor.orNull
        case 6 => limit.orNull
        case 8 => agentId.orNull
        case 9 => unread.orNull
      }
    }
    def getField(__field: _root_.scalapb.descriptors.FieldDescriptor): _root_.scalapb.descriptors.PValue = {
      require(__field.containingMessage eq companion.scalaDescriptor)
      (__field.number: @_root_.scala.unchecked) match {
        case 1 => ctx.map(_.toPMessage).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 3 => segmentationId.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 4 => query.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 5 => anchor.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 6 => limit.map(_root_.scalapb.descriptors.PInt).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 8 => agentId.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 9 => unread.map(_root_.scalapb.descriptors.PBoolean).getOrElse(_root_.scalapb.descriptors.PEmpty)
      }
    }
    def toProtoString: _root_.scala.Predef.String = _root_.scalapb.TextFormat.printToUnicodeString(this)
    def companion = user.UserSearchRequest
}

object UserSearchRequest extends scalapb.GeneratedMessageCompanion[user.UserSearchRequest] {
  implicit def messageCompanion: scalapb.GeneratedMessageCompanion[user.UserSearchRequest] = this
  def fromFieldsMap(__fieldsMap: scala.collection.immutable.Map[_root_.com.google.protobuf.Descriptors.FieldDescriptor, scala.Any]): user.UserSearchRequest = {
    require(__fieldsMap.keys.forall(_.getContainingType() == javaDescriptor), "FieldDescriptor does not match message type.")
    val __fields = javaDescriptor.getFields
    user.UserSearchRequest(
      __fieldsMap.get(__fields.get(0)).asInstanceOf[scala.Option[common.Context]],
      __fieldsMap.get(__fields.get(1)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(2)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(3)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(4)).asInstanceOf[scala.Option[_root_.scala.Int]],
      __fieldsMap.get(__fields.get(5)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(6)).asInstanceOf[scala.Option[_root_.scala.Boolean]]
    )
  }
  implicit def messageReads: _root_.scalapb.descriptors.Reads[user.UserSearchRequest] = _root_.scalapb.descriptors.Reads{
    case _root_.scalapb.descriptors.PMessage(__fieldsMap) =>
      require(__fieldsMap.keys.forall(_.containingMessage == scalaDescriptor), "FieldDescriptor does not match message type.")
      user.UserSearchRequest(
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(1).get).flatMap(_.as[scala.Option[common.Context]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(3).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(4).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(5).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(6).get).flatMap(_.as[scala.Option[_root_.scala.Int]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(8).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(9).get).flatMap(_.as[scala.Option[_root_.scala.Boolean]])
      )
    case _ => throw new RuntimeException("Expected PMessage")
  }
  def javaDescriptor: _root_.com.google.protobuf.Descriptors.Descriptor = UserProto.javaDescriptor.getMessageTypes.get(28)
  def scalaDescriptor: _root_.scalapb.descriptors.Descriptor = UserProto.scalaDescriptor.messages(28)
  def messageCompanionForFieldNumber(__number: _root_.scala.Int): _root_.scalapb.GeneratedMessageCompanion[_] = {
    var __out: _root_.scalapb.GeneratedMessageCompanion[_] = null
    (__number: @_root_.scala.unchecked) match {
      case 1 => __out = common.Context
    }
    __out
  }
  lazy val nestedMessagesCompanions: Seq[_root_.scalapb.GeneratedMessageCompanion[_]] = Seq.empty
  def enumCompanionForFieldNumber(__fieldNumber: _root_.scala.Int): _root_.scalapb.GeneratedEnumCompanion[_] = throw new MatchError(__fieldNumber)
  lazy val defaultInstance = user.UserSearchRequest(
  )
  implicit class UserSearchRequestLens[UpperPB](_l: _root_.scalapb.lenses.Lens[UpperPB, user.UserSearchRequest]) extends _root_.scalapb.lenses.ObjectLens[UpperPB, user.UserSearchRequest](_l) {
    def ctx: _root_.scalapb.lenses.Lens[UpperPB, common.Context] = field(_.getCtx)((c_, f_) => c_.copy(ctx = Option(f_)))
    def optionalCtx: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[common.Context]] = field(_.ctx)((c_, f_) => c_.copy(ctx = f_))
    def segmentationId: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getSegmentationId)((c_, f_) => c_.copy(segmentationId = Option(f_)))
    def optionalSegmentationId: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.segmentationId)((c_, f_) => c_.copy(segmentationId = f_))
    def query: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getQuery)((c_, f_) => c_.copy(query = Option(f_)))
    def optionalQuery: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.query)((c_, f_) => c_.copy(query = f_))
    def anchor: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getAnchor)((c_, f_) => c_.copy(anchor = Option(f_)))
    def optionalAnchor: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.anchor)((c_, f_) => c_.copy(anchor = f_))
    def limit: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Int] = field(_.getLimit)((c_, f_) => c_.copy(limit = Option(f_)))
    def optionalLimit: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Int]] = field(_.limit)((c_, f_) => c_.copy(limit = f_))
    def agentId: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getAgentId)((c_, f_) => c_.copy(agentId = Option(f_)))
    def optionalAgentId: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.agentId)((c_, f_) => c_.copy(agentId = f_))
    def unread: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Boolean] = field(_.getUnread)((c_, f_) => c_.copy(unread = Option(f_)))
    def optionalUnread: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Boolean]] = field(_.unread)((c_, f_) => c_.copy(unread = f_))
  }
  final val CTX_FIELD_NUMBER = 1
  final val SEGMENTATION_ID_FIELD_NUMBER = 3
  final val QUERY_FIELD_NUMBER = 4
  final val ANCHOR_FIELD_NUMBER = 5
  final val LIMIT_FIELD_NUMBER = 6
  final val AGENT_ID_FIELD_NUMBER = 8
  final val UNREAD_FIELD_NUMBER = 9
}
