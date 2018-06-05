// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO2

package api

@SerialVersionUID(0L)
final case class ListRequest(
    ctx: scala.Option[common.Context] = None,
    accountId: scala.Option[_root_.scala.Predef.String] = None,
    start: scala.Option[_root_.scala.Predef.String] = None,
    limit: scala.Option[_root_.scala.Int] = None
    ) extends scalapb.GeneratedMessage with scalapb.Message[ListRequest] with scalapb.lenses.Updatable[ListRequest] {
    @transient
    private[this] var __serializedSizeCachedValue: _root_.scala.Int = 0
    private[this] def __computeSerializedValue(): _root_.scala.Int = {
      var __size = 0
      if (ctx.isDefined) { __size += 1 + _root_.com.google.protobuf.CodedOutputStream.computeUInt32SizeNoTag(ctx.get.serializedSize) + ctx.get.serializedSize }
      if (accountId.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(2, accountId.get) }
      if (start.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(3, start.get) }
      if (limit.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeInt32Size(4, limit.get) }
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
      accountId.foreach { __v =>
        _output__.writeString(2, __v)
      };
      start.foreach { __v =>
        _output__.writeString(3, __v)
      };
      limit.foreach { __v =>
        _output__.writeInt32(4, __v)
      };
    }
    def mergeFrom(`_input__`: _root_.com.google.protobuf.CodedInputStream): api.ListRequest = {
      var __ctx = this.ctx
      var __accountId = this.accountId
      var __start = this.start
      var __limit = this.limit
      var _done__ = false
      while (!_done__) {
        val _tag__ = _input__.readTag()
        _tag__ match {
          case 0 => _done__ = true
          case 10 =>
            __ctx = Option(_root_.scalapb.LiteParser.readMessage(_input__, __ctx.getOrElse(common.Context.defaultInstance)))
          case 18 =>
            __accountId = Option(_input__.readString())
          case 26 =>
            __start = Option(_input__.readString())
          case 32 =>
            __limit = Option(_input__.readInt32())
          case tag => _input__.skipField(tag)
        }
      }
      api.ListRequest(
          ctx = __ctx,
          accountId = __accountId,
          start = __start,
          limit = __limit
      )
    }
    def getCtx: common.Context = ctx.getOrElse(common.Context.defaultInstance)
    def clearCtx: ListRequest = copy(ctx = None)
    def withCtx(__v: common.Context): ListRequest = copy(ctx = Option(__v))
    def getAccountId: _root_.scala.Predef.String = accountId.getOrElse("")
    def clearAccountId: ListRequest = copy(accountId = None)
    def withAccountId(__v: _root_.scala.Predef.String): ListRequest = copy(accountId = Option(__v))
    def getStart: _root_.scala.Predef.String = start.getOrElse("")
    def clearStart: ListRequest = copy(start = None)
    def withStart(__v: _root_.scala.Predef.String): ListRequest = copy(start = Option(__v))
    def getLimit: _root_.scala.Int = limit.getOrElse(0)
    def clearLimit: ListRequest = copy(limit = None)
    def withLimit(__v: _root_.scala.Int): ListRequest = copy(limit = Option(__v))
    def getFieldByNumber(__fieldNumber: _root_.scala.Int): scala.Any = {
      (__fieldNumber: @_root_.scala.unchecked) match {
        case 1 => ctx.orNull
        case 2 => accountId.orNull
        case 3 => start.orNull
        case 4 => limit.orNull
      }
    }
    def getField(__field: _root_.scalapb.descriptors.FieldDescriptor): _root_.scalapb.descriptors.PValue = {
      require(__field.containingMessage eq companion.scalaDescriptor)
      (__field.number: @_root_.scala.unchecked) match {
        case 1 => ctx.map(_.toPMessage).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 2 => accountId.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 3 => start.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 4 => limit.map(_root_.scalapb.descriptors.PInt).getOrElse(_root_.scalapb.descriptors.PEmpty)
      }
    }
    def toProtoString: _root_.scala.Predef.String = _root_.scalapb.TextFormat.printToUnicodeString(this)
    def companion = api.ListRequest
}

object ListRequest extends scalapb.GeneratedMessageCompanion[api.ListRequest] {
  implicit def messageCompanion: scalapb.GeneratedMessageCompanion[api.ListRequest] = this
  def fromFieldsMap(__fieldsMap: scala.collection.immutable.Map[_root_.com.google.protobuf.Descriptors.FieldDescriptor, scala.Any]): api.ListRequest = {
    require(__fieldsMap.keys.forall(_.getContainingType() == javaDescriptor), "FieldDescriptor does not match message type.")
    val __fields = javaDescriptor.getFields
    api.ListRequest(
      __fieldsMap.get(__fields.get(0)).asInstanceOf[scala.Option[common.Context]],
      __fieldsMap.get(__fields.get(1)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(2)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(3)).asInstanceOf[scala.Option[_root_.scala.Int]]
    )
  }
  implicit def messageReads: _root_.scalapb.descriptors.Reads[api.ListRequest] = _root_.scalapb.descriptors.Reads{
    case _root_.scalapb.descriptors.PMessage(__fieldsMap) =>
      require(__fieldsMap.keys.forall(_.containingMessage == scalaDescriptor), "FieldDescriptor does not match message type.")
      api.ListRequest(
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(1).get).flatMap(_.as[scala.Option[common.Context]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(2).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(3).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(4).get).flatMap(_.as[scala.Option[_root_.scala.Int]])
      )
    case _ => throw new RuntimeException("Expected PMessage")
  }
  def javaDescriptor: _root_.com.google.protobuf.Descriptors.Descriptor = ApiProto.javaDescriptor.getMessageTypes.get(7)
  def scalaDescriptor: _root_.scalapb.descriptors.Descriptor = ApiProto.scalaDescriptor.messages(7)
  def messageCompanionForFieldNumber(__number: _root_.scala.Int): _root_.scalapb.GeneratedMessageCompanion[_] = {
    var __out: _root_.scalapb.GeneratedMessageCompanion[_] = null
    (__number: @_root_.scala.unchecked) match {
      case 1 => __out = common.Context
    }
    __out
  }
  lazy val nestedMessagesCompanions: Seq[_root_.scalapb.GeneratedMessageCompanion[_]] = Seq.empty
  def enumCompanionForFieldNumber(__fieldNumber: _root_.scala.Int): _root_.scalapb.GeneratedEnumCompanion[_] = throw new MatchError(__fieldNumber)
  lazy val defaultInstance = api.ListRequest(
  )
  implicit class ListRequestLens[UpperPB](_l: _root_.scalapb.lenses.Lens[UpperPB, api.ListRequest]) extends _root_.scalapb.lenses.ObjectLens[UpperPB, api.ListRequest](_l) {
    def ctx: _root_.scalapb.lenses.Lens[UpperPB, common.Context] = field(_.getCtx)((c_, f_) => c_.copy(ctx = Option(f_)))
    def optionalCtx: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[common.Context]] = field(_.ctx)((c_, f_) => c_.copy(ctx = f_))
    def accountId: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getAccountId)((c_, f_) => c_.copy(accountId = Option(f_)))
    def optionalAccountId: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.accountId)((c_, f_) => c_.copy(accountId = f_))
    def start: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getStart)((c_, f_) => c_.copy(start = Option(f_)))
    def optionalStart: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.start)((c_, f_) => c_.copy(start = f_))
    def limit: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Int] = field(_.getLimit)((c_, f_) => c_.copy(limit = Option(f_)))
    def optionalLimit: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Int]] = field(_.limit)((c_, f_) => c_.copy(limit = f_))
  }
  final val CTX_FIELD_NUMBER = 1
  final val ACCOUNT_ID_FIELD_NUMBER = 2
  final val START_FIELD_NUMBER = 3
  final val LIMIT_FIELD_NUMBER = 4
}
