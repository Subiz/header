// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO3

package email

/** @param subject
  *  string to = 4;
  * @param attachments
  *   s3 links
  */
@SerialVersionUID(0L)
final case class Email(
    ctx: scala.Option[common.Context] = None,
    from: _root_.scala.Predef.String = "",
    subject: _root_.scala.Predef.String = "",
    body: _root_.scala.Predef.String = "",
    header: scala.collection.immutable.Map[_root_.scala.Predef.String, _root_.scala.Predef.String] = scala.collection.immutable.Map.empty,
    attachments: _root_.scala.collection.Seq[email.Attachment] = _root_.scala.collection.Seq.empty,
    to: _root_.scala.collection.Seq[_root_.scala.Predef.String] = _root_.scala.collection.Seq.empty,
    cc: _root_.scala.collection.Seq[_root_.scala.Predef.String] = _root_.scala.collection.Seq.empty,
    bcc: _root_.scala.collection.Seq[_root_.scala.Predef.String] = _root_.scala.collection.Seq.empty
    ) extends scalapb.GeneratedMessage with scalapb.Message[Email] with scalapb.lenses.Updatable[Email] {
    @transient
    private[this] var __serializedSizeCachedValue: _root_.scala.Int = 0
    private[this] def __computeSerializedValue(): _root_.scala.Int = {
      var __size = 0
      if (ctx.isDefined) { __size += 1 + _root_.com.google.protobuf.CodedOutputStream.computeUInt32SizeNoTag(ctx.get.serializedSize) + ctx.get.serializedSize }
      if (from != "") { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(3, from) }
      if (subject != "") { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(5, subject) }
      if (body != "") { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(6, body) }
      header.foreach(header => __size += 1 + _root_.com.google.protobuf.CodedOutputStream.computeUInt32SizeNoTag(email.Email._typemapper_header.toBase(header).serializedSize) + email.Email._typemapper_header.toBase(header).serializedSize)
      attachments.foreach(attachments => __size += 1 + _root_.com.google.protobuf.CodedOutputStream.computeUInt32SizeNoTag(attachments.serializedSize) + attachments.serializedSize)
      to.foreach(to => __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(9, to))
      cc.foreach(cc => __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(10, cc))
      bcc.foreach(bcc => __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(11, bcc))
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
        val __v = from
        if (__v != "") {
          _output__.writeString(3, __v)
        }
      };
      {
        val __v = subject
        if (__v != "") {
          _output__.writeString(5, __v)
        }
      };
      {
        val __v = body
        if (__v != "") {
          _output__.writeString(6, __v)
        }
      };
      header.foreach { __v =>
        _output__.writeTag(7, 2)
        _output__.writeUInt32NoTag(email.Email._typemapper_header.toBase(__v).serializedSize)
        email.Email._typemapper_header.toBase(__v).writeTo(_output__)
      };
      attachments.foreach { __v =>
        _output__.writeTag(8, 2)
        _output__.writeUInt32NoTag(__v.serializedSize)
        __v.writeTo(_output__)
      };
      to.foreach { __v =>
        _output__.writeString(9, __v)
      };
      cc.foreach { __v =>
        _output__.writeString(10, __v)
      };
      bcc.foreach { __v =>
        _output__.writeString(11, __v)
      };
    }
    def mergeFrom(`_input__`: _root_.com.google.protobuf.CodedInputStream): email.Email = {
      var __ctx = this.ctx
      var __from = this.from
      var __subject = this.subject
      var __body = this.body
      val __header = (scala.collection.immutable.Map.newBuilder[_root_.scala.Predef.String, _root_.scala.Predef.String] ++= this.header)
      val __attachments = (_root_.scala.collection.immutable.Vector.newBuilder[email.Attachment] ++= this.attachments)
      val __to = (_root_.scala.collection.immutable.Vector.newBuilder[_root_.scala.Predef.String] ++= this.to)
      val __cc = (_root_.scala.collection.immutable.Vector.newBuilder[_root_.scala.Predef.String] ++= this.cc)
      val __bcc = (_root_.scala.collection.immutable.Vector.newBuilder[_root_.scala.Predef.String] ++= this.bcc)
      var _done__ = false
      while (!_done__) {
        val _tag__ = _input__.readTag()
        _tag__ match {
          case 0 => _done__ = true
          case 10 =>
            __ctx = Option(_root_.scalapb.LiteParser.readMessage(_input__, __ctx.getOrElse(common.Context.defaultInstance)))
          case 26 =>
            __from = _input__.readString()
          case 42 =>
            __subject = _input__.readString()
          case 50 =>
            __body = _input__.readString()
          case 58 =>
            __header += email.Email._typemapper_header.toCustom(_root_.scalapb.LiteParser.readMessage(_input__, email.Email.HeaderEntry.defaultInstance))
          case 66 =>
            __attachments += _root_.scalapb.LiteParser.readMessage(_input__, email.Attachment.defaultInstance)
          case 74 =>
            __to += _input__.readString()
          case 82 =>
            __cc += _input__.readString()
          case 90 =>
            __bcc += _input__.readString()
          case tag => _input__.skipField(tag)
        }
      }
      email.Email(
          ctx = __ctx,
          from = __from,
          subject = __subject,
          body = __body,
          header = __header.result(),
          attachments = __attachments.result(),
          to = __to.result(),
          cc = __cc.result(),
          bcc = __bcc.result()
      )
    }
    def getCtx: common.Context = ctx.getOrElse(common.Context.defaultInstance)
    def clearCtx: Email = copy(ctx = None)
    def withCtx(__v: common.Context): Email = copy(ctx = Option(__v))
    def withFrom(__v: _root_.scala.Predef.String): Email = copy(from = __v)
    def withSubject(__v: _root_.scala.Predef.String): Email = copy(subject = __v)
    def withBody(__v: _root_.scala.Predef.String): Email = copy(body = __v)
    def clearHeader = copy(header = scala.collection.immutable.Map.empty)
    def addHeader(__vs: (_root_.scala.Predef.String, _root_.scala.Predef.String)*): Email = addAllHeader(__vs)
    def addAllHeader(__vs: TraversableOnce[(_root_.scala.Predef.String, _root_.scala.Predef.String)]): Email = copy(header = header ++ __vs)
    def withHeader(__v: scala.collection.immutable.Map[_root_.scala.Predef.String, _root_.scala.Predef.String]): Email = copy(header = __v)
    def clearAttachments = copy(attachments = _root_.scala.collection.Seq.empty)
    def addAttachments(__vs: email.Attachment*): Email = addAllAttachments(__vs)
    def addAllAttachments(__vs: TraversableOnce[email.Attachment]): Email = copy(attachments = attachments ++ __vs)
    def withAttachments(__v: _root_.scala.collection.Seq[email.Attachment]): Email = copy(attachments = __v)
    def clearTo = copy(to = _root_.scala.collection.Seq.empty)
    def addTo(__vs: _root_.scala.Predef.String*): Email = addAllTo(__vs)
    def addAllTo(__vs: TraversableOnce[_root_.scala.Predef.String]): Email = copy(to = to ++ __vs)
    def withTo(__v: _root_.scala.collection.Seq[_root_.scala.Predef.String]): Email = copy(to = __v)
    def clearCc = copy(cc = _root_.scala.collection.Seq.empty)
    def addCc(__vs: _root_.scala.Predef.String*): Email = addAllCc(__vs)
    def addAllCc(__vs: TraversableOnce[_root_.scala.Predef.String]): Email = copy(cc = cc ++ __vs)
    def withCc(__v: _root_.scala.collection.Seq[_root_.scala.Predef.String]): Email = copy(cc = __v)
    def clearBcc = copy(bcc = _root_.scala.collection.Seq.empty)
    def addBcc(__vs: _root_.scala.Predef.String*): Email = addAllBcc(__vs)
    def addAllBcc(__vs: TraversableOnce[_root_.scala.Predef.String]): Email = copy(bcc = bcc ++ __vs)
    def withBcc(__v: _root_.scala.collection.Seq[_root_.scala.Predef.String]): Email = copy(bcc = __v)
    def getFieldByNumber(__fieldNumber: _root_.scala.Int): scala.Any = {
      (__fieldNumber: @_root_.scala.unchecked) match {
        case 1 => ctx.orNull
        case 3 => {
          val __t = from
          if (__t != "") __t else null
        }
        case 5 => {
          val __t = subject
          if (__t != "") __t else null
        }
        case 6 => {
          val __t = body
          if (__t != "") __t else null
        }
        case 7 => header.map(email.Email._typemapper_header.toBase)(_root_.scala.collection.breakOut)
        case 8 => attachments
        case 9 => to
        case 10 => cc
        case 11 => bcc
      }
    }
    def getField(__field: _root_.scalapb.descriptors.FieldDescriptor): _root_.scalapb.descriptors.PValue = {
      require(__field.containingMessage eq companion.scalaDescriptor)
      (__field.number: @_root_.scala.unchecked) match {
        case 1 => ctx.map(_.toPMessage).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 3 => _root_.scalapb.descriptors.PString(from)
        case 5 => _root_.scalapb.descriptors.PString(subject)
        case 6 => _root_.scalapb.descriptors.PString(body)
        case 7 => _root_.scalapb.descriptors.PRepeated(header.map(email.Email._typemapper_header.toBase(_).toPMessage)(_root_.scala.collection.breakOut))
        case 8 => _root_.scalapb.descriptors.PRepeated(attachments.map(_.toPMessage)(_root_.scala.collection.breakOut))
        case 9 => _root_.scalapb.descriptors.PRepeated(to.map(_root_.scalapb.descriptors.PString)(_root_.scala.collection.breakOut))
        case 10 => _root_.scalapb.descriptors.PRepeated(cc.map(_root_.scalapb.descriptors.PString)(_root_.scala.collection.breakOut))
        case 11 => _root_.scalapb.descriptors.PRepeated(bcc.map(_root_.scalapb.descriptors.PString)(_root_.scala.collection.breakOut))
      }
    }
    def toProtoString: _root_.scala.Predef.String = _root_.scalapb.TextFormat.printToUnicodeString(this)
    def companion = email.Email
}

object Email extends scalapb.GeneratedMessageCompanion[email.Email] {
  implicit def messageCompanion: scalapb.GeneratedMessageCompanion[email.Email] = this
  def fromFieldsMap(__fieldsMap: scala.collection.immutable.Map[_root_.com.google.protobuf.Descriptors.FieldDescriptor, scala.Any]): email.Email = {
    require(__fieldsMap.keys.forall(_.getContainingType() == javaDescriptor), "FieldDescriptor does not match message type.")
    val __fields = javaDescriptor.getFields
    email.Email(
      __fieldsMap.get(__fields.get(0)).asInstanceOf[scala.Option[common.Context]],
      __fieldsMap.getOrElse(__fields.get(1), "").asInstanceOf[_root_.scala.Predef.String],
      __fieldsMap.getOrElse(__fields.get(2), "").asInstanceOf[_root_.scala.Predef.String],
      __fieldsMap.getOrElse(__fields.get(3), "").asInstanceOf[_root_.scala.Predef.String],
      __fieldsMap.getOrElse(__fields.get(4), Nil).asInstanceOf[_root_.scala.collection.Seq[email.Email.HeaderEntry]].map(email.Email._typemapper_header.toCustom)(_root_.scala.collection.breakOut),
      __fieldsMap.getOrElse(__fields.get(5), Nil).asInstanceOf[_root_.scala.collection.Seq[email.Attachment]],
      __fieldsMap.getOrElse(__fields.get(6), Nil).asInstanceOf[_root_.scala.collection.Seq[_root_.scala.Predef.String]],
      __fieldsMap.getOrElse(__fields.get(7), Nil).asInstanceOf[_root_.scala.collection.Seq[_root_.scala.Predef.String]],
      __fieldsMap.getOrElse(__fields.get(8), Nil).asInstanceOf[_root_.scala.collection.Seq[_root_.scala.Predef.String]]
    )
  }
  implicit def messageReads: _root_.scalapb.descriptors.Reads[email.Email] = _root_.scalapb.descriptors.Reads{
    case _root_.scalapb.descriptors.PMessage(__fieldsMap) =>
      require(__fieldsMap.keys.forall(_.containingMessage == scalaDescriptor), "FieldDescriptor does not match message type.")
      email.Email(
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(1).get).flatMap(_.as[scala.Option[common.Context]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(3).get).map(_.as[_root_.scala.Predef.String]).getOrElse(""),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(5).get).map(_.as[_root_.scala.Predef.String]).getOrElse(""),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(6).get).map(_.as[_root_.scala.Predef.String]).getOrElse(""),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(7).get).map(_.as[_root_.scala.collection.Seq[email.Email.HeaderEntry]]).getOrElse(_root_.scala.collection.Seq.empty).map(email.Email._typemapper_header.toCustom)(_root_.scala.collection.breakOut),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(8).get).map(_.as[_root_.scala.collection.Seq[email.Attachment]]).getOrElse(_root_.scala.collection.Seq.empty),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(9).get).map(_.as[_root_.scala.collection.Seq[_root_.scala.Predef.String]]).getOrElse(_root_.scala.collection.Seq.empty),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(10).get).map(_.as[_root_.scala.collection.Seq[_root_.scala.Predef.String]]).getOrElse(_root_.scala.collection.Seq.empty),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(11).get).map(_.as[_root_.scala.collection.Seq[_root_.scala.Predef.String]]).getOrElse(_root_.scala.collection.Seq.empty)
      )
    case _ => throw new RuntimeException("Expected PMessage")
  }
  def javaDescriptor: _root_.com.google.protobuf.Descriptors.Descriptor = EmailProto.javaDescriptor.getMessageTypes.get(0)
  def scalaDescriptor: _root_.scalapb.descriptors.Descriptor = EmailProto.scalaDescriptor.messages(0)
  def messageCompanionForFieldNumber(__number: _root_.scala.Int): _root_.scalapb.GeneratedMessageCompanion[_] = {
    var __out: _root_.scalapb.GeneratedMessageCompanion[_] = null
    (__number: @_root_.scala.unchecked) match {
      case 1 => __out = common.Context
      case 7 => __out = email.Email.HeaderEntry
      case 8 => __out = email.Attachment
    }
    __out
  }
  lazy val nestedMessagesCompanions: Seq[_root_.scalapb.GeneratedMessageCompanion[_]] = Seq[_root_.scalapb.GeneratedMessageCompanion[_]](
    _root_.email.Email.HeaderEntry
  )
  def enumCompanionForFieldNumber(__fieldNumber: _root_.scala.Int): _root_.scalapb.GeneratedEnumCompanion[_] = throw new MatchError(__fieldNumber)
  lazy val defaultInstance = email.Email(
  )
  @SerialVersionUID(0L)
  final case class HeaderEntry(
      key: _root_.scala.Predef.String = "",
      value: _root_.scala.Predef.String = ""
      ) extends scalapb.GeneratedMessage with scalapb.Message[HeaderEntry] with scalapb.lenses.Updatable[HeaderEntry] {
      @transient
      private[this] var __serializedSizeCachedValue: _root_.scala.Int = 0
      private[this] def __computeSerializedValue(): _root_.scala.Int = {
        var __size = 0
        if (key != "") { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(1, key) }
        if (value != "") { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(2, value) }
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
        {
          val __v = key
          if (__v != "") {
            _output__.writeString(1, __v)
          }
        };
        {
          val __v = value
          if (__v != "") {
            _output__.writeString(2, __v)
          }
        };
      }
      def mergeFrom(`_input__`: _root_.com.google.protobuf.CodedInputStream): email.Email.HeaderEntry = {
        var __key = this.key
        var __value = this.value
        var _done__ = false
        while (!_done__) {
          val _tag__ = _input__.readTag()
          _tag__ match {
            case 0 => _done__ = true
            case 10 =>
              __key = _input__.readString()
            case 18 =>
              __value = _input__.readString()
            case tag => _input__.skipField(tag)
          }
        }
        email.Email.HeaderEntry(
            key = __key,
            value = __value
        )
      }
      def withKey(__v: _root_.scala.Predef.String): HeaderEntry = copy(key = __v)
      def withValue(__v: _root_.scala.Predef.String): HeaderEntry = copy(value = __v)
      def getFieldByNumber(__fieldNumber: _root_.scala.Int): scala.Any = {
        (__fieldNumber: @_root_.scala.unchecked) match {
          case 1 => {
            val __t = key
            if (__t != "") __t else null
          }
          case 2 => {
            val __t = value
            if (__t != "") __t else null
          }
        }
      }
      def getField(__field: _root_.scalapb.descriptors.FieldDescriptor): _root_.scalapb.descriptors.PValue = {
        require(__field.containingMessage eq companion.scalaDescriptor)
        (__field.number: @_root_.scala.unchecked) match {
          case 1 => _root_.scalapb.descriptors.PString(key)
          case 2 => _root_.scalapb.descriptors.PString(value)
        }
      }
      def toProtoString: _root_.scala.Predef.String = _root_.scalapb.TextFormat.printToUnicodeString(this)
      def companion = email.Email.HeaderEntry
  }
  
  object HeaderEntry extends scalapb.GeneratedMessageCompanion[email.Email.HeaderEntry] {
    implicit def messageCompanion: scalapb.GeneratedMessageCompanion[email.Email.HeaderEntry] = this
    def fromFieldsMap(__fieldsMap: scala.collection.immutable.Map[_root_.com.google.protobuf.Descriptors.FieldDescriptor, scala.Any]): email.Email.HeaderEntry = {
      require(__fieldsMap.keys.forall(_.getContainingType() == javaDescriptor), "FieldDescriptor does not match message type.")
      val __fields = javaDescriptor.getFields
      email.Email.HeaderEntry(
        __fieldsMap.getOrElse(__fields.get(0), "").asInstanceOf[_root_.scala.Predef.String],
        __fieldsMap.getOrElse(__fields.get(1), "").asInstanceOf[_root_.scala.Predef.String]
      )
    }
    implicit def messageReads: _root_.scalapb.descriptors.Reads[email.Email.HeaderEntry] = _root_.scalapb.descriptors.Reads{
      case _root_.scalapb.descriptors.PMessage(__fieldsMap) =>
        require(__fieldsMap.keys.forall(_.containingMessage == scalaDescriptor), "FieldDescriptor does not match message type.")
        email.Email.HeaderEntry(
          __fieldsMap.get(scalaDescriptor.findFieldByNumber(1).get).map(_.as[_root_.scala.Predef.String]).getOrElse(""),
          __fieldsMap.get(scalaDescriptor.findFieldByNumber(2).get).map(_.as[_root_.scala.Predef.String]).getOrElse("")
        )
      case _ => throw new RuntimeException("Expected PMessage")
    }
    def javaDescriptor: _root_.com.google.protobuf.Descriptors.Descriptor = email.Email.javaDescriptor.getNestedTypes.get(0)
    def scalaDescriptor: _root_.scalapb.descriptors.Descriptor = email.Email.scalaDescriptor.nestedMessages(0)
    def messageCompanionForFieldNumber(__number: _root_.scala.Int): _root_.scalapb.GeneratedMessageCompanion[_] = throw new MatchError(__number)
    lazy val nestedMessagesCompanions: Seq[_root_.scalapb.GeneratedMessageCompanion[_]] = Seq.empty
    def enumCompanionForFieldNumber(__fieldNumber: _root_.scala.Int): _root_.scalapb.GeneratedEnumCompanion[_] = throw new MatchError(__fieldNumber)
    lazy val defaultInstance = email.Email.HeaderEntry(
    )
    implicit class HeaderEntryLens[UpperPB](_l: _root_.scalapb.lenses.Lens[UpperPB, email.Email.HeaderEntry]) extends _root_.scalapb.lenses.ObjectLens[UpperPB, email.Email.HeaderEntry](_l) {
      def key: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.key)((c_, f_) => c_.copy(key = f_))
      def value: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.value)((c_, f_) => c_.copy(value = f_))
    }
    final val KEY_FIELD_NUMBER = 1
    final val VALUE_FIELD_NUMBER = 2
    implicit val keyValueMapper: _root_.scalapb.TypeMapper[email.Email.HeaderEntry, (_root_.scala.Predef.String, _root_.scala.Predef.String)] =
      _root_.scalapb.TypeMapper[email.Email.HeaderEntry, (_root_.scala.Predef.String, _root_.scala.Predef.String)](__m => (__m.key, __m.value))(__p => email.Email.HeaderEntry(__p._1, __p._2))
  }
  
  implicit class EmailLens[UpperPB](_l: _root_.scalapb.lenses.Lens[UpperPB, email.Email]) extends _root_.scalapb.lenses.ObjectLens[UpperPB, email.Email](_l) {
    def ctx: _root_.scalapb.lenses.Lens[UpperPB, common.Context] = field(_.getCtx)((c_, f_) => c_.copy(ctx = Option(f_)))
    def optionalCtx: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[common.Context]] = field(_.ctx)((c_, f_) => c_.copy(ctx = f_))
    def from: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.from)((c_, f_) => c_.copy(from = f_))
    def subject: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.subject)((c_, f_) => c_.copy(subject = f_))
    def body: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.body)((c_, f_) => c_.copy(body = f_))
    def header: _root_.scalapb.lenses.Lens[UpperPB, scala.collection.immutable.Map[_root_.scala.Predef.String, _root_.scala.Predef.String]] = field(_.header)((c_, f_) => c_.copy(header = f_))
    def attachments: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.collection.Seq[email.Attachment]] = field(_.attachments)((c_, f_) => c_.copy(attachments = f_))
    def to: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.collection.Seq[_root_.scala.Predef.String]] = field(_.to)((c_, f_) => c_.copy(to = f_))
    def cc: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.collection.Seq[_root_.scala.Predef.String]] = field(_.cc)((c_, f_) => c_.copy(cc = f_))
    def bcc: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.collection.Seq[_root_.scala.Predef.String]] = field(_.bcc)((c_, f_) => c_.copy(bcc = f_))
  }
  final val CTX_FIELD_NUMBER = 1
  final val FROM_FIELD_NUMBER = 3
  final val SUBJECT_FIELD_NUMBER = 5
  final val BODY_FIELD_NUMBER = 6
  final val HEADER_FIELD_NUMBER = 7
  final val ATTACHMENTS_FIELD_NUMBER = 8
  final val TO_FIELD_NUMBER = 9
  final val CC_FIELD_NUMBER = 10
  final val BCC_FIELD_NUMBER = 11
  @transient
  private val _typemapper_header: _root_.scalapb.TypeMapper[email.Email.HeaderEntry, (_root_.scala.Predef.String, _root_.scala.Predef.String)] = implicitly[_root_.scalapb.TypeMapper[email.Email.HeaderEntry, (_root_.scala.Predef.String, _root_.scala.Predef.String)]]
}