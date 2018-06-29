// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO3

package file

@SerialVersionUID(0L)
final case class File(
    ctx: scala.Option[common.Context] = None,
    accountId: _root_.scala.Predef.String = "",
    name: _root_.scala.Predef.String = "",
    `type`: _root_.scala.Predef.String = "",
    size: _root_.scala.Long = 0L,
    md5: _root_.scala.Predef.String = "",
    description: _root_.scala.Predef.String = "",
    created: _root_.scala.Long = 0L,
    url: _root_.scala.Predef.String = "",
    creator: _root_.scala.Predef.String = "",
    id: _root_.scala.Predef.String = ""
    ) extends scalapb.GeneratedMessage with scalapb.Message[File] with scalapb.lenses.Updatable[File] {
    @transient
    private[this] var __serializedSizeCachedValue: _root_.scala.Int = 0
    private[this] def __computeSerializedValue(): _root_.scala.Int = {
      var __size = 0
      if (ctx.isDefined) { __size += 1 + _root_.com.google.protobuf.CodedOutputStream.computeUInt32SizeNoTag(ctx.get.serializedSize) + ctx.get.serializedSize }
      if (accountId != "") { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(2, accountId) }
      if (name != "") { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(3, name) }
      if (`type` != "") { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(4, `type`) }
      if (size != 0L) { __size += _root_.com.google.protobuf.CodedOutputStream.computeInt64Size(5, size) }
      if (md5 != "") { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(6, md5) }
      if (description != "") { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(10, description) }
      if (created != 0L) { __size += _root_.com.google.protobuf.CodedOutputStream.computeInt64Size(7, created) }
      if (url != "") { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(8, url) }
      if (creator != "") { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(9, creator) }
      if (id != "") { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(11, id) }
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
      {
        val __v = accountId
        if (__v != "") {
          _output__.writeString(2, __v)
        }
      };
      {
        val __v = name
        if (__v != "") {
          _output__.writeString(3, __v)
        }
      };
      {
        val __v = `type`
        if (__v != "") {
          _output__.writeString(4, __v)
        }
      };
      {
        val __v = size
        if (__v != 0L) {
          _output__.writeInt64(5, __v)
        }
      };
      {
        val __v = md5
        if (__v != "") {
          _output__.writeString(6, __v)
        }
      };
      {
        val __v = created
        if (__v != 0L) {
          _output__.writeInt64(7, __v)
        }
      };
      {
        val __v = url
        if (__v != "") {
          _output__.writeString(8, __v)
        }
      };
      {
        val __v = creator
        if (__v != "") {
          _output__.writeString(9, __v)
        }
      };
      {
        val __v = description
        if (__v != "") {
          _output__.writeString(10, __v)
        }
      };
      {
        val __v = id
        if (__v != "") {
          _output__.writeString(11, __v)
        }
      };
    }
    def mergeFrom(`_input__`: _root_.com.google.protobuf.CodedInputStream): file.File = {
      var __ctx = this.ctx
      var __accountId = this.accountId
      var __name = this.name
      var __type = this.`type`
      var __size = this.size
      var __md5 = this.md5
      var __description = this.description
      var __created = this.created
      var __url = this.url
      var __creator = this.creator
      var __id = this.id
      var _done__ = false
      while (!_done__) {
        val _tag__ = _input__.readTag()
        _tag__ match {
          case 0 => _done__ = true
          case 10 =>
            __ctx = Option(_root_.scalapb.LiteParser.readMessage(_input__, __ctx.getOrElse(common.Context.defaultInstance)))
          case 18 =>
            __accountId = _input__.readString()
          case 26 =>
            __name = _input__.readString()
          case 34 =>
            __type = _input__.readString()
          case 40 =>
            __size = _input__.readInt64()
          case 50 =>
            __md5 = _input__.readString()
          case 82 =>
            __description = _input__.readString()
          case 56 =>
            __created = _input__.readInt64()
          case 66 =>
            __url = _input__.readString()
          case 74 =>
            __creator = _input__.readString()
          case 90 =>
            __id = _input__.readString()
          case tag => _input__.skipField(tag)
        }
      }
      file.File(
          ctx = __ctx,
          accountId = __accountId,
          name = __name,
          `type` = __type,
          size = __size,
          md5 = __md5,
          description = __description,
          created = __created,
          url = __url,
          creator = __creator,
          id = __id
      )
    }
    def getCtx: common.Context = ctx.getOrElse(common.Context.defaultInstance)
    def clearCtx: File = copy(ctx = None)
    def withCtx(__v: common.Context): File = copy(ctx = Option(__v))
    def withAccountId(__v: _root_.scala.Predef.String): File = copy(accountId = __v)
    def withName(__v: _root_.scala.Predef.String): File = copy(name = __v)
    def withType(__v: _root_.scala.Predef.String): File = copy(`type` = __v)
    def withSize(__v: _root_.scala.Long): File = copy(size = __v)
    def withMd5(__v: _root_.scala.Predef.String): File = copy(md5 = __v)
    def withDescription(__v: _root_.scala.Predef.String): File = copy(description = __v)
    def withCreated(__v: _root_.scala.Long): File = copy(created = __v)
    def withUrl(__v: _root_.scala.Predef.String): File = copy(url = __v)
    def withCreator(__v: _root_.scala.Predef.String): File = copy(creator = __v)
    def withId(__v: _root_.scala.Predef.String): File = copy(id = __v)
    def getFieldByNumber(__fieldNumber: _root_.scala.Int): scala.Any = {
      (__fieldNumber: @_root_.scala.unchecked) match {
        case 1 => ctx.orNull
        case 2 => {
          val __t = accountId
          if (__t != "") __t else null
        }
        case 3 => {
          val __t = name
          if (__t != "") __t else null
        }
        case 4 => {
          val __t = `type`
          if (__t != "") __t else null
        }
        case 5 => {
          val __t = size
          if (__t != 0L) __t else null
        }
        case 6 => {
          val __t = md5
          if (__t != "") __t else null
        }
        case 10 => {
          val __t = description
          if (__t != "") __t else null
        }
        case 7 => {
          val __t = created
          if (__t != 0L) __t else null
        }
        case 8 => {
          val __t = url
          if (__t != "") __t else null
        }
        case 9 => {
          val __t = creator
          if (__t != "") __t else null
        }
        case 11 => {
          val __t = id
          if (__t != "") __t else null
        }
      }
    }
    def getField(__field: _root_.scalapb.descriptors.FieldDescriptor): _root_.scalapb.descriptors.PValue = {
      require(__field.containingMessage eq companion.scalaDescriptor)
      (__field.number: @_root_.scala.unchecked) match {
        case 1 => ctx.map(_.toPMessage).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 2 => _root_.scalapb.descriptors.PString(accountId)
        case 3 => _root_.scalapb.descriptors.PString(name)
        case 4 => _root_.scalapb.descriptors.PString(`type`)
        case 5 => _root_.scalapb.descriptors.PLong(size)
        case 6 => _root_.scalapb.descriptors.PString(md5)
        case 10 => _root_.scalapb.descriptors.PString(description)
        case 7 => _root_.scalapb.descriptors.PLong(created)
        case 8 => _root_.scalapb.descriptors.PString(url)
        case 9 => _root_.scalapb.descriptors.PString(creator)
        case 11 => _root_.scalapb.descriptors.PString(id)
      }
    }
    def toProtoString: _root_.scala.Predef.String = _root_.scalapb.TextFormat.printToUnicodeString(this)
    def companion = file.File
}

object File extends scalapb.GeneratedMessageCompanion[file.File] {
  implicit def messageCompanion: scalapb.GeneratedMessageCompanion[file.File] = this
  def fromFieldsMap(__fieldsMap: scala.collection.immutable.Map[_root_.com.google.protobuf.Descriptors.FieldDescriptor, scala.Any]): file.File = {
    require(__fieldsMap.keys.forall(_.getContainingType() == javaDescriptor), "FieldDescriptor does not match message type.")
    val __fields = javaDescriptor.getFields
    file.File(
      __fieldsMap.get(__fields.get(0)).asInstanceOf[scala.Option[common.Context]],
      __fieldsMap.getOrElse(__fields.get(1), "").asInstanceOf[_root_.scala.Predef.String],
      __fieldsMap.getOrElse(__fields.get(2), "").asInstanceOf[_root_.scala.Predef.String],
      __fieldsMap.getOrElse(__fields.get(3), "").asInstanceOf[_root_.scala.Predef.String],
      __fieldsMap.getOrElse(__fields.get(4), 0L).asInstanceOf[_root_.scala.Long],
      __fieldsMap.getOrElse(__fields.get(5), "").asInstanceOf[_root_.scala.Predef.String],
      __fieldsMap.getOrElse(__fields.get(6), "").asInstanceOf[_root_.scala.Predef.String],
      __fieldsMap.getOrElse(__fields.get(7), 0L).asInstanceOf[_root_.scala.Long],
      __fieldsMap.getOrElse(__fields.get(8), "").asInstanceOf[_root_.scala.Predef.String],
      __fieldsMap.getOrElse(__fields.get(9), "").asInstanceOf[_root_.scala.Predef.String],
      __fieldsMap.getOrElse(__fields.get(10), "").asInstanceOf[_root_.scala.Predef.String]
    )
  }
  implicit def messageReads: _root_.scalapb.descriptors.Reads[file.File] = _root_.scalapb.descriptors.Reads{
    case _root_.scalapb.descriptors.PMessage(__fieldsMap) =>
      require(__fieldsMap.keys.forall(_.containingMessage == scalaDescriptor), "FieldDescriptor does not match message type.")
      file.File(
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(1).get).flatMap(_.as[scala.Option[common.Context]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(2).get).map(_.as[_root_.scala.Predef.String]).getOrElse(""),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(3).get).map(_.as[_root_.scala.Predef.String]).getOrElse(""),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(4).get).map(_.as[_root_.scala.Predef.String]).getOrElse(""),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(5).get).map(_.as[_root_.scala.Long]).getOrElse(0L),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(6).get).map(_.as[_root_.scala.Predef.String]).getOrElse(""),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(10).get).map(_.as[_root_.scala.Predef.String]).getOrElse(""),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(7).get).map(_.as[_root_.scala.Long]).getOrElse(0L),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(8).get).map(_.as[_root_.scala.Predef.String]).getOrElse(""),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(9).get).map(_.as[_root_.scala.Predef.String]).getOrElse(""),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(11).get).map(_.as[_root_.scala.Predef.String]).getOrElse("")
      )
    case _ => throw new RuntimeException("Expected PMessage")
  }
  def javaDescriptor: _root_.com.google.protobuf.Descriptors.Descriptor = FileProto.javaDescriptor.getMessageTypes.get(4)
  def scalaDescriptor: _root_.scalapb.descriptors.Descriptor = FileProto.scalaDescriptor.messages(4)
  def messageCompanionForFieldNumber(__number: _root_.scala.Int): _root_.scalapb.GeneratedMessageCompanion[_] = {
    var __out: _root_.scalapb.GeneratedMessageCompanion[_] = null
    (__number: @_root_.scala.unchecked) match {
      case 1 => __out = common.Context
    }
    __out
  }
  lazy val nestedMessagesCompanions: Seq[_root_.scalapb.GeneratedMessageCompanion[_]] = Seq.empty
  def enumCompanionForFieldNumber(__fieldNumber: _root_.scala.Int): _root_.scalapb.GeneratedEnumCompanion[_] = throw new MatchError(__fieldNumber)
  lazy val defaultInstance = file.File(
  )
  implicit class FileLens[UpperPB](_l: _root_.scalapb.lenses.Lens[UpperPB, file.File]) extends _root_.scalapb.lenses.ObjectLens[UpperPB, file.File](_l) {
    def ctx: _root_.scalapb.lenses.Lens[UpperPB, common.Context] = field(_.getCtx)((c_, f_) => c_.copy(ctx = Option(f_)))
    def optionalCtx: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[common.Context]] = field(_.ctx)((c_, f_) => c_.copy(ctx = f_))
    def accountId: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.accountId)((c_, f_) => c_.copy(accountId = f_))
    def name: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.name)((c_, f_) => c_.copy(name = f_))
    def `type`: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.`type`)((c_, f_) => c_.copy(`type` = f_))
    def size: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Long] = field(_.size)((c_, f_) => c_.copy(size = f_))
    def md5: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.md5)((c_, f_) => c_.copy(md5 = f_))
    def description: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.description)((c_, f_) => c_.copy(description = f_))
    def created: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Long] = field(_.created)((c_, f_) => c_.copy(created = f_))
    def url: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.url)((c_, f_) => c_.copy(url = f_))
    def creator: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.creator)((c_, f_) => c_.copy(creator = f_))
    def id: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.id)((c_, f_) => c_.copy(id = f_))
  }
  final val CTX_FIELD_NUMBER = 1
  final val ACCOUNT_ID_FIELD_NUMBER = 2
  final val NAME_FIELD_NUMBER = 3
  final val TYPE_FIELD_NUMBER = 4
  final val SIZE_FIELD_NUMBER = 5
  final val MD5_FIELD_NUMBER = 6
  final val DESCRIPTION_FIELD_NUMBER = 10
  final val CREATED_FIELD_NUMBER = 7
  final val URL_FIELD_NUMBER = 8
  final val CREATOR_FIELD_NUMBER = 9
  final val ID_FIELD_NUMBER = 11
}