// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO2

package notibox

@SerialVersionUID(0L)
final case class ReadNotification(
    ctx: scala.Option[common.Context] = None,
    box: scala.Option[_root_.scala.Predef.String] = None,
    topics: _root_.scala.collection.Seq[_root_.scala.Predef.String] = _root_.scala.collection.Seq.empty,
    read: scala.Option[_root_.scala.Boolean] = None
    ) extends scalapb.GeneratedMessage with scalapb.Message[ReadNotification] with scalapb.lenses.Updatable[ReadNotification] {
    @transient
    private[this] var __serializedSizeCachedValue: _root_.scala.Int = 0
    private[this] def __computeSerializedValue(): _root_.scala.Int = {
      var __size = 0
      if (ctx.isDefined) { __size += 1 + _root_.com.google.protobuf.CodedOutputStream.computeUInt32SizeNoTag(ctx.get.serializedSize) + ctx.get.serializedSize }
      if (box.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(2, box.get) }
      topics.foreach(topics => __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(3, topics))
      if (read.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeBoolSize(4, read.get) }
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
      box.foreach { __v =>
        _output__.writeString(2, __v)
      };
      topics.foreach { __v =>
        _output__.writeString(3, __v)
      };
      read.foreach { __v =>
        _output__.writeBool(4, __v)
      };
    }
    def mergeFrom(`_input__`: _root_.com.google.protobuf.CodedInputStream): notibox.ReadNotification = {
      var __ctx = this.ctx
      var __box = this.box
      val __topics = (_root_.scala.collection.immutable.Vector.newBuilder[_root_.scala.Predef.String] ++= this.topics)
      var __read = this.read
      var _done__ = false
      while (!_done__) {
        val _tag__ = _input__.readTag()
        _tag__ match {
          case 0 => _done__ = true
          case 10 =>
            __ctx = Option(_root_.scalapb.LiteParser.readMessage(_input__, __ctx.getOrElse(common.Context.defaultInstance)))
          case 18 =>
            __box = Option(_input__.readString())
          case 26 =>
            __topics += _input__.readString()
          case 32 =>
            __read = Option(_input__.readBool())
          case tag => _input__.skipField(tag)
        }
      }
      notibox.ReadNotification(
          ctx = __ctx,
          box = __box,
          topics = __topics.result(),
          read = __read
      )
    }
    def getCtx: common.Context = ctx.getOrElse(common.Context.defaultInstance)
    def clearCtx: ReadNotification = copy(ctx = None)
    def withCtx(__v: common.Context): ReadNotification = copy(ctx = Option(__v))
    def getBox: _root_.scala.Predef.String = box.getOrElse("")
    def clearBox: ReadNotification = copy(box = None)
    def withBox(__v: _root_.scala.Predef.String): ReadNotification = copy(box = Option(__v))
    def clearTopics = copy(topics = _root_.scala.collection.Seq.empty)
    def addTopics(__vs: _root_.scala.Predef.String*): ReadNotification = addAllTopics(__vs)
    def addAllTopics(__vs: TraversableOnce[_root_.scala.Predef.String]): ReadNotification = copy(topics = topics ++ __vs)
    def withTopics(__v: _root_.scala.collection.Seq[_root_.scala.Predef.String]): ReadNotification = copy(topics = __v)
    def getRead: _root_.scala.Boolean = read.getOrElse(false)
    def clearRead: ReadNotification = copy(read = None)
    def withRead(__v: _root_.scala.Boolean): ReadNotification = copy(read = Option(__v))
    def getFieldByNumber(__fieldNumber: _root_.scala.Int): scala.Any = {
      (__fieldNumber: @_root_.scala.unchecked) match {
        case 1 => ctx.orNull
        case 2 => box.orNull
        case 3 => topics
        case 4 => read.orNull
      }
    }
    def getField(__field: _root_.scalapb.descriptors.FieldDescriptor): _root_.scalapb.descriptors.PValue = {
      require(__field.containingMessage eq companion.scalaDescriptor)
      (__field.number: @_root_.scala.unchecked) match {
        case 1 => ctx.map(_.toPMessage).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 2 => box.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 3 => _root_.scalapb.descriptors.PRepeated(topics.map(_root_.scalapb.descriptors.PString)(_root_.scala.collection.breakOut))
        case 4 => read.map(_root_.scalapb.descriptors.PBoolean).getOrElse(_root_.scalapb.descriptors.PEmpty)
      }
    }
    def toProtoString: _root_.scala.Predef.String = _root_.scalapb.TextFormat.printToUnicodeString(this)
    def companion = notibox.ReadNotification
}

object ReadNotification extends scalapb.GeneratedMessageCompanion[notibox.ReadNotification] {
  implicit def messageCompanion: scalapb.GeneratedMessageCompanion[notibox.ReadNotification] = this
  def fromFieldsMap(__fieldsMap: scala.collection.immutable.Map[_root_.com.google.protobuf.Descriptors.FieldDescriptor, scala.Any]): notibox.ReadNotification = {
    require(__fieldsMap.keys.forall(_.getContainingType() == javaDescriptor), "FieldDescriptor does not match message type.")
    val __fields = javaDescriptor.getFields
    notibox.ReadNotification(
      __fieldsMap.get(__fields.get(0)).asInstanceOf[scala.Option[common.Context]],
      __fieldsMap.get(__fields.get(1)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.getOrElse(__fields.get(2), Nil).asInstanceOf[_root_.scala.collection.Seq[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(3)).asInstanceOf[scala.Option[_root_.scala.Boolean]]
    )
  }
  implicit def messageReads: _root_.scalapb.descriptors.Reads[notibox.ReadNotification] = _root_.scalapb.descriptors.Reads{
    case _root_.scalapb.descriptors.PMessage(__fieldsMap) =>
      require(__fieldsMap.keys.forall(_.containingMessage == scalaDescriptor), "FieldDescriptor does not match message type.")
      notibox.ReadNotification(
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(1).get).flatMap(_.as[scala.Option[common.Context]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(2).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(3).get).map(_.as[_root_.scala.collection.Seq[_root_.scala.Predef.String]]).getOrElse(_root_.scala.collection.Seq.empty),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(4).get).flatMap(_.as[scala.Option[_root_.scala.Boolean]])
      )
    case _ => throw new RuntimeException("Expected PMessage")
  }
  def javaDescriptor: _root_.com.google.protobuf.Descriptors.Descriptor = NotiboxProto.javaDescriptor.getMessageTypes.get(2)
  def scalaDescriptor: _root_.scalapb.descriptors.Descriptor = NotiboxProto.scalaDescriptor.messages(2)
  def messageCompanionForFieldNumber(__number: _root_.scala.Int): _root_.scalapb.GeneratedMessageCompanion[_] = {
    var __out: _root_.scalapb.GeneratedMessageCompanion[_] = null
    (__number: @_root_.scala.unchecked) match {
      case 1 => __out = common.Context
    }
    __out
  }
  lazy val nestedMessagesCompanions: Seq[_root_.scalapb.GeneratedMessageCompanion[_]] = Seq.empty
  def enumCompanionForFieldNumber(__fieldNumber: _root_.scala.Int): _root_.scalapb.GeneratedEnumCompanion[_] = throw new MatchError(__fieldNumber)
  lazy val defaultInstance = notibox.ReadNotification(
  )
  implicit class ReadNotificationLens[UpperPB](_l: _root_.scalapb.lenses.Lens[UpperPB, notibox.ReadNotification]) extends _root_.scalapb.lenses.ObjectLens[UpperPB, notibox.ReadNotification](_l) {
    def ctx: _root_.scalapb.lenses.Lens[UpperPB, common.Context] = field(_.getCtx)((c_, f_) => c_.copy(ctx = Option(f_)))
    def optionalCtx: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[common.Context]] = field(_.ctx)((c_, f_) => c_.copy(ctx = f_))
    def box: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getBox)((c_, f_) => c_.copy(box = Option(f_)))
    def optionalBox: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.box)((c_, f_) => c_.copy(box = f_))
    def topics: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.collection.Seq[_root_.scala.Predef.String]] = field(_.topics)((c_, f_) => c_.copy(topics = f_))
    def read: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Boolean] = field(_.getRead)((c_, f_) => c_.copy(read = Option(f_)))
    def optionalRead: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Boolean]] = field(_.read)((c_, f_) => c_.copy(read = f_))
  }
  final val CTX_FIELD_NUMBER = 1
  final val BOX_FIELD_NUMBER = 2
  final val TOPICS_FIELD_NUMBER = 3
  final val READ_FIELD_NUMBER = 4
}
