// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO2

package user

@SerialVersionUID(0L)
final case class AttributeData(
    ctx: scala.Option[common.Context] = None,
    accountId: scala.Option[_root_.scala.Predef.String] = None,
    userId: scala.Option[_root_.scala.Predef.String] = None,
    key: scala.Option[_root_.scala.Predef.String] = None,
    value: scala.Option[_root_.scala.Predef.String] = None,
    state: scala.Option[_root_.scala.Predef.String] = None,
    created: scala.Option[_root_.scala.Long] = None,
    modified: scala.Option[_root_.scala.Long] = None
    ) extends scalapb.GeneratedMessage with scalapb.Message[AttributeData] with scalapb.lenses.Updatable[AttributeData] {
    @transient
    private[this] var __serializedSizeCachedValue: _root_.scala.Int = 0
    private[this] def __computeSerializedValue(): _root_.scala.Int = {
      var __size = 0
      if (ctx.isDefined) { __size += 1 + _root_.com.google.protobuf.CodedOutputStream.computeUInt32SizeNoTag(ctx.get.serializedSize) + ctx.get.serializedSize }
      if (accountId.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(2, accountId.get) }
      if (userId.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(3, userId.get) }
      if (key.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(4, key.get) }
      if (value.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(5, value.get) }
      if (state.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(6, state.get) }
      if (created.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeInt64Size(7, created.get) }
      if (modified.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeInt64Size(8, modified.get) }
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
      userId.foreach { __v =>
        _output__.writeString(3, __v)
      };
      key.foreach { __v =>
        _output__.writeString(4, __v)
      };
      value.foreach { __v =>
        _output__.writeString(5, __v)
      };
      state.foreach { __v =>
        _output__.writeString(6, __v)
      };
      created.foreach { __v =>
        _output__.writeInt64(7, __v)
      };
      modified.foreach { __v =>
        _output__.writeInt64(8, __v)
      };
    }
    def mergeFrom(`_input__`: _root_.com.google.protobuf.CodedInputStream): user.AttributeData = {
      var __ctx = this.ctx
      var __accountId = this.accountId
      var __userId = this.userId
      var __key = this.key
      var __value = this.value
      var __state = this.state
      var __created = this.created
      var __modified = this.modified
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
            __userId = Option(_input__.readString())
          case 34 =>
            __key = Option(_input__.readString())
          case 42 =>
            __value = Option(_input__.readString())
          case 50 =>
            __state = Option(_input__.readString())
          case 56 =>
            __created = Option(_input__.readInt64())
          case 64 =>
            __modified = Option(_input__.readInt64())
          case tag => _input__.skipField(tag)
        }
      }
      user.AttributeData(
          ctx = __ctx,
          accountId = __accountId,
          userId = __userId,
          key = __key,
          value = __value,
          state = __state,
          created = __created,
          modified = __modified
      )
    }
    def getCtx: common.Context = ctx.getOrElse(common.Context.defaultInstance)
    def clearCtx: AttributeData = copy(ctx = None)
    def withCtx(__v: common.Context): AttributeData = copy(ctx = Option(__v))
    def getAccountId: _root_.scala.Predef.String = accountId.getOrElse("")
    def clearAccountId: AttributeData = copy(accountId = None)
    def withAccountId(__v: _root_.scala.Predef.String): AttributeData = copy(accountId = Option(__v))
    def getUserId: _root_.scala.Predef.String = userId.getOrElse("")
    def clearUserId: AttributeData = copy(userId = None)
    def withUserId(__v: _root_.scala.Predef.String): AttributeData = copy(userId = Option(__v))
    def getKey: _root_.scala.Predef.String = key.getOrElse("")
    def clearKey: AttributeData = copy(key = None)
    def withKey(__v: _root_.scala.Predef.String): AttributeData = copy(key = Option(__v))
    def getValue: _root_.scala.Predef.String = value.getOrElse("")
    def clearValue: AttributeData = copy(value = None)
    def withValue(__v: _root_.scala.Predef.String): AttributeData = copy(value = Option(__v))
    def getState: _root_.scala.Predef.String = state.getOrElse("")
    def clearState: AttributeData = copy(state = None)
    def withState(__v: _root_.scala.Predef.String): AttributeData = copy(state = Option(__v))
    def getCreated: _root_.scala.Long = created.getOrElse(0L)
    def clearCreated: AttributeData = copy(created = None)
    def withCreated(__v: _root_.scala.Long): AttributeData = copy(created = Option(__v))
    def getModified: _root_.scala.Long = modified.getOrElse(0L)
    def clearModified: AttributeData = copy(modified = None)
    def withModified(__v: _root_.scala.Long): AttributeData = copy(modified = Option(__v))
    def getFieldByNumber(__fieldNumber: _root_.scala.Int): scala.Any = {
      (__fieldNumber: @_root_.scala.unchecked) match {
        case 1 => ctx.orNull
        case 2 => accountId.orNull
        case 3 => userId.orNull
        case 4 => key.orNull
        case 5 => value.orNull
        case 6 => state.orNull
        case 7 => created.orNull
        case 8 => modified.orNull
      }
    }
    def getField(__field: _root_.scalapb.descriptors.FieldDescriptor): _root_.scalapb.descriptors.PValue = {
      require(__field.containingMessage eq companion.scalaDescriptor)
      (__field.number: @_root_.scala.unchecked) match {
        case 1 => ctx.map(_.toPMessage).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 2 => accountId.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 3 => userId.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 4 => key.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 5 => value.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 6 => state.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 7 => created.map(_root_.scalapb.descriptors.PLong).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 8 => modified.map(_root_.scalapb.descriptors.PLong).getOrElse(_root_.scalapb.descriptors.PEmpty)
      }
    }
    def toProtoString: _root_.scala.Predef.String = _root_.scalapb.TextFormat.printToUnicodeString(this)
    def companion = user.AttributeData
}

object AttributeData extends scalapb.GeneratedMessageCompanion[user.AttributeData] {
  implicit def messageCompanion: scalapb.GeneratedMessageCompanion[user.AttributeData] = this
  def fromFieldsMap(__fieldsMap: scala.collection.immutable.Map[_root_.com.google.protobuf.Descriptors.FieldDescriptor, scala.Any]): user.AttributeData = {
    require(__fieldsMap.keys.forall(_.getContainingType() == javaDescriptor), "FieldDescriptor does not match message type.")
    val __fields = javaDescriptor.getFields
    user.AttributeData(
      __fieldsMap.get(__fields.get(0)).asInstanceOf[scala.Option[common.Context]],
      __fieldsMap.get(__fields.get(1)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(2)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(3)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(4)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(5)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(6)).asInstanceOf[scala.Option[_root_.scala.Long]],
      __fieldsMap.get(__fields.get(7)).asInstanceOf[scala.Option[_root_.scala.Long]]
    )
  }
  implicit def messageReads: _root_.scalapb.descriptors.Reads[user.AttributeData] = _root_.scalapb.descriptors.Reads{
    case _root_.scalapb.descriptors.PMessage(__fieldsMap) =>
      require(__fieldsMap.keys.forall(_.containingMessage == scalaDescriptor), "FieldDescriptor does not match message type.")
      user.AttributeData(
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(1).get).flatMap(_.as[scala.Option[common.Context]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(2).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(3).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(4).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(5).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(6).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(7).get).flatMap(_.as[scala.Option[_root_.scala.Long]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(8).get).flatMap(_.as[scala.Option[_root_.scala.Long]])
      )
    case _ => throw new RuntimeException("Expected PMessage")
  }
  def javaDescriptor: _root_.com.google.protobuf.Descriptors.Descriptor = UserProto.javaDescriptor.getMessageTypes.get(6)
  def scalaDescriptor: _root_.scalapb.descriptors.Descriptor = UserProto.scalaDescriptor.messages(6)
  def messageCompanionForFieldNumber(__number: _root_.scala.Int): _root_.scalapb.GeneratedMessageCompanion[_] = {
    var __out: _root_.scalapb.GeneratedMessageCompanion[_] = null
    (__number: @_root_.scala.unchecked) match {
      case 1 => __out = common.Context
    }
    __out
  }
  lazy val nestedMessagesCompanions: Seq[_root_.scalapb.GeneratedMessageCompanion[_]] = Seq.empty
  def enumCompanionForFieldNumber(__fieldNumber: _root_.scala.Int): _root_.scalapb.GeneratedEnumCompanion[_] = throw new MatchError(__fieldNumber)
  lazy val defaultInstance = user.AttributeData(
  )
  implicit class AttributeDataLens[UpperPB](_l: _root_.scalapb.lenses.Lens[UpperPB, user.AttributeData]) extends _root_.scalapb.lenses.ObjectLens[UpperPB, user.AttributeData](_l) {
    def ctx: _root_.scalapb.lenses.Lens[UpperPB, common.Context] = field(_.getCtx)((c_, f_) => c_.copy(ctx = Option(f_)))
    def optionalCtx: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[common.Context]] = field(_.ctx)((c_, f_) => c_.copy(ctx = f_))
    def accountId: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getAccountId)((c_, f_) => c_.copy(accountId = Option(f_)))
    def optionalAccountId: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.accountId)((c_, f_) => c_.copy(accountId = f_))
    def userId: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getUserId)((c_, f_) => c_.copy(userId = Option(f_)))
    def optionalUserId: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.userId)((c_, f_) => c_.copy(userId = f_))
    def key: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getKey)((c_, f_) => c_.copy(key = Option(f_)))
    def optionalKey: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.key)((c_, f_) => c_.copy(key = f_))
    def value: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getValue)((c_, f_) => c_.copy(value = Option(f_)))
    def optionalValue: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.value)((c_, f_) => c_.copy(value = f_))
    def state: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getState)((c_, f_) => c_.copy(state = Option(f_)))
    def optionalState: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.state)((c_, f_) => c_.copy(state = f_))
    def created: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Long] = field(_.getCreated)((c_, f_) => c_.copy(created = Option(f_)))
    def optionalCreated: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Long]] = field(_.created)((c_, f_) => c_.copy(created = f_))
    def modified: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Long] = field(_.getModified)((c_, f_) => c_.copy(modified = Option(f_)))
    def optionalModified: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Long]] = field(_.modified)((c_, f_) => c_.copy(modified = f_))
  }
  final val CTX_FIELD_NUMBER = 1
  final val ACCOUNT_ID_FIELD_NUMBER = 2
  final val USER_ID_FIELD_NUMBER = 3
  final val KEY_FIELD_NUMBER = 4
  final val VALUE_FIELD_NUMBER = 5
  final val STATE_FIELD_NUMBER = 6
  final val CREATED_FIELD_NUMBER = 7
  final val MODIFIED_FIELD_NUMBER = 8
}
