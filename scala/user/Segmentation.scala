// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO2

package user

/** @param ran
  *   repeated Condition conditions = 6;
  *  repeated SegmentCondition query = 5; // suffix annotation
  *   a * b + (c + d) * e  ==&gt;  + * a b * + c d e
  *   time start run
  */
@SerialVersionUID(0L)
final case class Segmentation(
    ctx: scala.Option[common.Context] = None,
    accountId: scala.Option[_root_.scala.Predef.String] = None,
    id: scala.Option[_root_.scala.Predef.String] = None,
    name: scala.Option[_root_.scala.Predef.String] = None,
    description: scala.Option[_root_.scala.Predef.String] = None,
    userCount: scala.Option[_root_.scala.Long] = None,
    ran: scala.Option[_root_.scala.Long] = None,
    startedFrom: scala.Option[_root_.scala.Long] = None,
    created: scala.Option[_root_.scala.Long] = None,
    modified: scala.Option[_root_.scala.Long] = None,
    state: scala.Option[_root_.scala.Predef.String] = None,
    condition: scala.Option[user.SegmentCondition] = None,
    currentCursor: scala.Option[_root_.scala.Predef.String] = None
    ) extends scalapb.GeneratedMessage with scalapb.Message[Segmentation] with scalapb.lenses.Updatable[Segmentation] {
    @transient
    private[this] var __serializedSizeCachedValue: _root_.scala.Int = 0
    private[this] def __computeSerializedValue(): _root_.scala.Int = {
      var __size = 0
      if (ctx.isDefined) { __size += 1 + _root_.com.google.protobuf.CodedOutputStream.computeUInt32SizeNoTag(ctx.get.serializedSize) + ctx.get.serializedSize }
      if (accountId.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(2, accountId.get) }
      if (id.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(5, id.get) }
      if (name.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(3, name.get) }
      if (description.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(11, description.get) }
      if (userCount.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeInt64Size(4, userCount.get) }
      if (ran.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeInt64Size(8, ran.get) }
      if (startedFrom.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeInt64Size(7, startedFrom.get) }
      if (created.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeInt64Size(9, created.get) }
      if (modified.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeInt64Size(10, modified.get) }
      if (state.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(12, state.get) }
      if (condition.isDefined) { __size += 2 + _root_.com.google.protobuf.CodedOutputStream.computeUInt32SizeNoTag(condition.get.serializedSize) + condition.get.serializedSize }
      if (currentCursor.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(18, currentCursor.get) }
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
      name.foreach { __v =>
        _output__.writeString(3, __v)
      };
      userCount.foreach { __v =>
        _output__.writeInt64(4, __v)
      };
      id.foreach { __v =>
        _output__.writeString(5, __v)
      };
      startedFrom.foreach { __v =>
        _output__.writeInt64(7, __v)
      };
      ran.foreach { __v =>
        _output__.writeInt64(8, __v)
      };
      created.foreach { __v =>
        _output__.writeInt64(9, __v)
      };
      modified.foreach { __v =>
        _output__.writeInt64(10, __v)
      };
      description.foreach { __v =>
        _output__.writeString(11, __v)
      };
      state.foreach { __v =>
        _output__.writeString(12, __v)
      };
      currentCursor.foreach { __v =>
        _output__.writeString(18, __v)
      };
      condition.foreach { __v =>
        _output__.writeTag(19, 2)
        _output__.writeUInt32NoTag(__v.serializedSize)
        __v.writeTo(_output__)
      };
    }
    def mergeFrom(`_input__`: _root_.com.google.protobuf.CodedInputStream): user.Segmentation = {
      var __ctx = this.ctx
      var __accountId = this.accountId
      var __id = this.id
      var __name = this.name
      var __description = this.description
      var __userCount = this.userCount
      var __ran = this.ran
      var __startedFrom = this.startedFrom
      var __created = this.created
      var __modified = this.modified
      var __state = this.state
      var __condition = this.condition
      var __currentCursor = this.currentCursor
      var _done__ = false
      while (!_done__) {
        val _tag__ = _input__.readTag()
        _tag__ match {
          case 0 => _done__ = true
          case 10 =>
            __ctx = Option(_root_.scalapb.LiteParser.readMessage(_input__, __ctx.getOrElse(common.Context.defaultInstance)))
          case 18 =>
            __accountId = Option(_input__.readString())
          case 42 =>
            __id = Option(_input__.readString())
          case 26 =>
            __name = Option(_input__.readString())
          case 90 =>
            __description = Option(_input__.readString())
          case 32 =>
            __userCount = Option(_input__.readInt64())
          case 64 =>
            __ran = Option(_input__.readInt64())
          case 56 =>
            __startedFrom = Option(_input__.readInt64())
          case 72 =>
            __created = Option(_input__.readInt64())
          case 80 =>
            __modified = Option(_input__.readInt64())
          case 98 =>
            __state = Option(_input__.readString())
          case 154 =>
            __condition = Option(_root_.scalapb.LiteParser.readMessage(_input__, __condition.getOrElse(user.SegmentCondition.defaultInstance)))
          case 146 =>
            __currentCursor = Option(_input__.readString())
          case tag => _input__.skipField(tag)
        }
      }
      user.Segmentation(
          ctx = __ctx,
          accountId = __accountId,
          id = __id,
          name = __name,
          description = __description,
          userCount = __userCount,
          ran = __ran,
          startedFrom = __startedFrom,
          created = __created,
          modified = __modified,
          state = __state,
          condition = __condition,
          currentCursor = __currentCursor
      )
    }
    def getCtx: common.Context = ctx.getOrElse(common.Context.defaultInstance)
    def clearCtx: Segmentation = copy(ctx = None)
    def withCtx(__v: common.Context): Segmentation = copy(ctx = Option(__v))
    def getAccountId: _root_.scala.Predef.String = accountId.getOrElse("")
    def clearAccountId: Segmentation = copy(accountId = None)
    def withAccountId(__v: _root_.scala.Predef.String): Segmentation = copy(accountId = Option(__v))
    def getId: _root_.scala.Predef.String = id.getOrElse("")
    def clearId: Segmentation = copy(id = None)
    def withId(__v: _root_.scala.Predef.String): Segmentation = copy(id = Option(__v))
    def getName: _root_.scala.Predef.String = name.getOrElse("")
    def clearName: Segmentation = copy(name = None)
    def withName(__v: _root_.scala.Predef.String): Segmentation = copy(name = Option(__v))
    def getDescription: _root_.scala.Predef.String = description.getOrElse("")
    def clearDescription: Segmentation = copy(description = None)
    def withDescription(__v: _root_.scala.Predef.String): Segmentation = copy(description = Option(__v))
    def getUserCount: _root_.scala.Long = userCount.getOrElse(0L)
    def clearUserCount: Segmentation = copy(userCount = None)
    def withUserCount(__v: _root_.scala.Long): Segmentation = copy(userCount = Option(__v))
    def getRan: _root_.scala.Long = ran.getOrElse(0L)
    def clearRan: Segmentation = copy(ran = None)
    def withRan(__v: _root_.scala.Long): Segmentation = copy(ran = Option(__v))
    def getStartedFrom: _root_.scala.Long = startedFrom.getOrElse(0L)
    def clearStartedFrom: Segmentation = copy(startedFrom = None)
    def withStartedFrom(__v: _root_.scala.Long): Segmentation = copy(startedFrom = Option(__v))
    def getCreated: _root_.scala.Long = created.getOrElse(0L)
    def clearCreated: Segmentation = copy(created = None)
    def withCreated(__v: _root_.scala.Long): Segmentation = copy(created = Option(__v))
    def getModified: _root_.scala.Long = modified.getOrElse(0L)
    def clearModified: Segmentation = copy(modified = None)
    def withModified(__v: _root_.scala.Long): Segmentation = copy(modified = Option(__v))
    def getState: _root_.scala.Predef.String = state.getOrElse("")
    def clearState: Segmentation = copy(state = None)
    def withState(__v: _root_.scala.Predef.String): Segmentation = copy(state = Option(__v))
    def getCondition: user.SegmentCondition = condition.getOrElse(user.SegmentCondition.defaultInstance)
    def clearCondition: Segmentation = copy(condition = None)
    def withCondition(__v: user.SegmentCondition): Segmentation = copy(condition = Option(__v))
    def getCurrentCursor: _root_.scala.Predef.String = currentCursor.getOrElse("")
    def clearCurrentCursor: Segmentation = copy(currentCursor = None)
    def withCurrentCursor(__v: _root_.scala.Predef.String): Segmentation = copy(currentCursor = Option(__v))
    def getFieldByNumber(__fieldNumber: _root_.scala.Int): scala.Any = {
      (__fieldNumber: @_root_.scala.unchecked) match {
        case 1 => ctx.orNull
        case 2 => accountId.orNull
        case 5 => id.orNull
        case 3 => name.orNull
        case 11 => description.orNull
        case 4 => userCount.orNull
        case 8 => ran.orNull
        case 7 => startedFrom.orNull
        case 9 => created.orNull
        case 10 => modified.orNull
        case 12 => state.orNull
        case 19 => condition.orNull
        case 18 => currentCursor.orNull
      }
    }
    def getField(__field: _root_.scalapb.descriptors.FieldDescriptor): _root_.scalapb.descriptors.PValue = {
      require(__field.containingMessage eq companion.scalaDescriptor)
      (__field.number: @_root_.scala.unchecked) match {
        case 1 => ctx.map(_.toPMessage).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 2 => accountId.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 5 => id.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 3 => name.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 11 => description.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 4 => userCount.map(_root_.scalapb.descriptors.PLong).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 8 => ran.map(_root_.scalapb.descriptors.PLong).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 7 => startedFrom.map(_root_.scalapb.descriptors.PLong).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 9 => created.map(_root_.scalapb.descriptors.PLong).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 10 => modified.map(_root_.scalapb.descriptors.PLong).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 12 => state.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 19 => condition.map(_.toPMessage).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 18 => currentCursor.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
      }
    }
    def toProtoString: _root_.scala.Predef.String = _root_.scalapb.TextFormat.printToUnicodeString(this)
    def companion = user.Segmentation
}

object Segmentation extends scalapb.GeneratedMessageCompanion[user.Segmentation] {
  implicit def messageCompanion: scalapb.GeneratedMessageCompanion[user.Segmentation] = this
  def fromFieldsMap(__fieldsMap: scala.collection.immutable.Map[_root_.com.google.protobuf.Descriptors.FieldDescriptor, scala.Any]): user.Segmentation = {
    require(__fieldsMap.keys.forall(_.getContainingType() == javaDescriptor), "FieldDescriptor does not match message type.")
    val __fields = javaDescriptor.getFields
    user.Segmentation(
      __fieldsMap.get(__fields.get(0)).asInstanceOf[scala.Option[common.Context]],
      __fieldsMap.get(__fields.get(1)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(2)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(3)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(4)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(5)).asInstanceOf[scala.Option[_root_.scala.Long]],
      __fieldsMap.get(__fields.get(6)).asInstanceOf[scala.Option[_root_.scala.Long]],
      __fieldsMap.get(__fields.get(7)).asInstanceOf[scala.Option[_root_.scala.Long]],
      __fieldsMap.get(__fields.get(8)).asInstanceOf[scala.Option[_root_.scala.Long]],
      __fieldsMap.get(__fields.get(9)).asInstanceOf[scala.Option[_root_.scala.Long]],
      __fieldsMap.get(__fields.get(10)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(11)).asInstanceOf[scala.Option[user.SegmentCondition]],
      __fieldsMap.get(__fields.get(12)).asInstanceOf[scala.Option[_root_.scala.Predef.String]]
    )
  }
  implicit def messageReads: _root_.scalapb.descriptors.Reads[user.Segmentation] = _root_.scalapb.descriptors.Reads{
    case _root_.scalapb.descriptors.PMessage(__fieldsMap) =>
      require(__fieldsMap.keys.forall(_.containingMessage == scalaDescriptor), "FieldDescriptor does not match message type.")
      user.Segmentation(
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(1).get).flatMap(_.as[scala.Option[common.Context]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(2).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(5).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(3).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(11).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(4).get).flatMap(_.as[scala.Option[_root_.scala.Long]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(8).get).flatMap(_.as[scala.Option[_root_.scala.Long]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(7).get).flatMap(_.as[scala.Option[_root_.scala.Long]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(9).get).flatMap(_.as[scala.Option[_root_.scala.Long]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(10).get).flatMap(_.as[scala.Option[_root_.scala.Long]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(12).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(19).get).flatMap(_.as[scala.Option[user.SegmentCondition]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(18).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]])
      )
    case _ => throw new RuntimeException("Expected PMessage")
  }
  def javaDescriptor: _root_.com.google.protobuf.Descriptors.Descriptor = UserProto.javaDescriptor.getMessageTypes.get(23)
  def scalaDescriptor: _root_.scalapb.descriptors.Descriptor = UserProto.scalaDescriptor.messages(23)
  def messageCompanionForFieldNumber(__number: _root_.scala.Int): _root_.scalapb.GeneratedMessageCompanion[_] = {
    var __out: _root_.scalapb.GeneratedMessageCompanion[_] = null
    (__number: @_root_.scala.unchecked) match {
      case 1 => __out = common.Context
      case 19 => __out = user.SegmentCondition
    }
    __out
  }
  lazy val nestedMessagesCompanions: Seq[_root_.scalapb.GeneratedMessageCompanion[_]] = Seq.empty
  def enumCompanionForFieldNumber(__fieldNumber: _root_.scala.Int): _root_.scalapb.GeneratedEnumCompanion[_] = throw new MatchError(__fieldNumber)
  lazy val defaultInstance = user.Segmentation(
  )
  sealed trait State extends _root_.scalapb.GeneratedEnum {
    type EnumType = State
    def isactive: _root_.scala.Boolean = false
    def isinactive: _root_.scala.Boolean = false
    def companion: _root_.scalapb.GeneratedEnumCompanion[State] = user.Segmentation.State
  }
  
  object State extends _root_.scalapb.GeneratedEnumCompanion[State] {
    implicit def enumCompanion: _root_.scalapb.GeneratedEnumCompanion[State] = this
    @SerialVersionUID(0L)
    case object active extends State {
      val value = 0
      val index = 0
      val name = "active"
      override def isactive: _root_.scala.Boolean = true
    }
    
    @SerialVersionUID(0L)
    case object inactive extends State {
      val value = 1
      val index = 1
      val name = "inactive"
      override def isinactive: _root_.scala.Boolean = true
    }
    
    @SerialVersionUID(0L)
    final case class Unrecognized(value: _root_.scala.Int) extends State with _root_.scalapb.UnrecognizedEnum
    
    lazy val values = scala.collection.Seq(active, inactive)
    def fromValue(value: _root_.scala.Int): State = value match {
      case 0 => active
      case 1 => inactive
      case __other => Unrecognized(__other)
    }
    def javaDescriptor: _root_.com.google.protobuf.Descriptors.EnumDescriptor = user.Segmentation.javaDescriptor.getEnumTypes.get(0)
    def scalaDescriptor: _root_.scalapb.descriptors.EnumDescriptor = user.Segmentation.scalaDescriptor.enums(0)
  }
  implicit class SegmentationLens[UpperPB](_l: _root_.scalapb.lenses.Lens[UpperPB, user.Segmentation]) extends _root_.scalapb.lenses.ObjectLens[UpperPB, user.Segmentation](_l) {
    def ctx: _root_.scalapb.lenses.Lens[UpperPB, common.Context] = field(_.getCtx)((c_, f_) => c_.copy(ctx = Option(f_)))
    def optionalCtx: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[common.Context]] = field(_.ctx)((c_, f_) => c_.copy(ctx = f_))
    def accountId: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getAccountId)((c_, f_) => c_.copy(accountId = Option(f_)))
    def optionalAccountId: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.accountId)((c_, f_) => c_.copy(accountId = f_))
    def id: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getId)((c_, f_) => c_.copy(id = Option(f_)))
    def optionalId: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.id)((c_, f_) => c_.copy(id = f_))
    def name: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getName)((c_, f_) => c_.copy(name = Option(f_)))
    def optionalName: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.name)((c_, f_) => c_.copy(name = f_))
    def description: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getDescription)((c_, f_) => c_.copy(description = Option(f_)))
    def optionalDescription: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.description)((c_, f_) => c_.copy(description = f_))
    def userCount: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Long] = field(_.getUserCount)((c_, f_) => c_.copy(userCount = Option(f_)))
    def optionalUserCount: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Long]] = field(_.userCount)((c_, f_) => c_.copy(userCount = f_))
    def ran: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Long] = field(_.getRan)((c_, f_) => c_.copy(ran = Option(f_)))
    def optionalRan: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Long]] = field(_.ran)((c_, f_) => c_.copy(ran = f_))
    def startedFrom: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Long] = field(_.getStartedFrom)((c_, f_) => c_.copy(startedFrom = Option(f_)))
    def optionalStartedFrom: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Long]] = field(_.startedFrom)((c_, f_) => c_.copy(startedFrom = f_))
    def created: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Long] = field(_.getCreated)((c_, f_) => c_.copy(created = Option(f_)))
    def optionalCreated: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Long]] = field(_.created)((c_, f_) => c_.copy(created = f_))
    def modified: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Long] = field(_.getModified)((c_, f_) => c_.copy(modified = Option(f_)))
    def optionalModified: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Long]] = field(_.modified)((c_, f_) => c_.copy(modified = f_))
    def state: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getState)((c_, f_) => c_.copy(state = Option(f_)))
    def optionalState: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.state)((c_, f_) => c_.copy(state = f_))
    def condition: _root_.scalapb.lenses.Lens[UpperPB, user.SegmentCondition] = field(_.getCondition)((c_, f_) => c_.copy(condition = Option(f_)))
    def optionalCondition: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[user.SegmentCondition]] = field(_.condition)((c_, f_) => c_.copy(condition = f_))
    def currentCursor: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getCurrentCursor)((c_, f_) => c_.copy(currentCursor = Option(f_)))
    def optionalCurrentCursor: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.currentCursor)((c_, f_) => c_.copy(currentCursor = f_))
  }
  final val CTX_FIELD_NUMBER = 1
  final val ACCOUNT_ID_FIELD_NUMBER = 2
  final val ID_FIELD_NUMBER = 5
  final val NAME_FIELD_NUMBER = 3
  final val DESCRIPTION_FIELD_NUMBER = 11
  final val USER_COUNT_FIELD_NUMBER = 4
  final val RAN_FIELD_NUMBER = 8
  final val STARTED_FROM_FIELD_NUMBER = 7
  final val CREATED_FIELD_NUMBER = 9
  final val MODIFIED_FIELD_NUMBER = 10
  final val STATE_FIELD_NUMBER = 12
  final val CONDITION_FIELD_NUMBER = 19
  final val CURRENT_CURSOR_FIELD_NUMBER = 18
}