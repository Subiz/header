// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO2

package conversation

/** @param mimetype
  *   file
  * @param length
  *  optional string data = 7; // for custom data
  * @param size
  *   byte
  * @param elements
  *   generic
  * @param title
  *   preview
  * @param askInfoAnswer
  *  repeated string inputtype = 17;
  */
@SerialVersionUID(0L)
final case class Attachment(
    `type`: scala.Option[_root_.scala.Predef.String] = None,
    mimetype: scala.Option[_root_.scala.Predef.String] = None,
    url: scala.Option[_root_.scala.Predef.String] = None,
    thumbnailUrl: scala.Option[_root_.scala.Predef.String] = None,
    name: scala.Option[_root_.scala.Predef.String] = None,
    description: scala.Option[_root_.scala.Predef.String] = None,
    length: scala.Option[_root_.scala.Int] = None,
    size: scala.Option[_root_.scala.Int] = None,
    elements: _root_.scala.collection.Seq[conversation.GenericElementTemplate] = _root_.scala.collection.Seq.empty,
    title: scala.Option[_root_.scala.Predef.String] = None,
    color: scala.Option[_root_.scala.Predef.String] = None,
    pretext: scala.Option[_root_.scala.Predef.String] = None,
    buttons: _root_.scala.collection.Seq[conversation.Button] = _root_.scala.collection.Seq.empty,
    askInfo: scala.Option[conversation.AskInfomation] = None,
    askInfoAnswer: scala.Option[conversation.AskInfomationAnswer] = None,
    form: scala.Option[conversation.Form] = None,
    formSubmit: scala.Option[conversation.FormSubmit] = None
    ) extends scalapb.GeneratedMessage with scalapb.Message[Attachment] with scalapb.lenses.Updatable[Attachment] {
    @transient
    private[this] var __serializedSizeCachedValue: _root_.scala.Int = 0
    private[this] def __computeSerializedValue(): _root_.scala.Int = {
      var __size = 0
      if (`type`.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(12, `type`.get) }
      if (mimetype.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(2, mimetype.get) }
      if (url.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(3, url.get) }
      if (thumbnailUrl.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(4, thumbnailUrl.get) }
      if (name.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(5, name.get) }
      if (description.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(6, description.get) }
      if (length.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeInt32Size(15, length.get) }
      if (size.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeInt32Size(13, size.get) }
      elements.foreach(elements => __size += 1 + _root_.com.google.protobuf.CodedOutputStream.computeUInt32SizeNoTag(elements.serializedSize) + elements.serializedSize)
      if (title.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(9, title.get) }
      if (color.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(10, color.get) }
      if (pretext.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(11, pretext.get) }
      buttons.foreach(buttons => __size += 2 + _root_.com.google.protobuf.CodedOutputStream.computeUInt32SizeNoTag(buttons.serializedSize) + buttons.serializedSize)
      if (askInfo.isDefined) { __size += 2 + _root_.com.google.protobuf.CodedOutputStream.computeUInt32SizeNoTag(askInfo.get.serializedSize) + askInfo.get.serializedSize }
      if (askInfoAnswer.isDefined) { __size += 2 + _root_.com.google.protobuf.CodedOutputStream.computeUInt32SizeNoTag(askInfoAnswer.get.serializedSize) + askInfoAnswer.get.serializedSize }
      if (form.isDefined) { __size += 2 + _root_.com.google.protobuf.CodedOutputStream.computeUInt32SizeNoTag(form.get.serializedSize) + form.get.serializedSize }
      if (formSubmit.isDefined) { __size += 2 + _root_.com.google.protobuf.CodedOutputStream.computeUInt32SizeNoTag(formSubmit.get.serializedSize) + formSubmit.get.serializedSize }
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
      mimetype.foreach { __v =>
        _output__.writeString(2, __v)
      };
      url.foreach { __v =>
        _output__.writeString(3, __v)
      };
      thumbnailUrl.foreach { __v =>
        _output__.writeString(4, __v)
      };
      name.foreach { __v =>
        _output__.writeString(5, __v)
      };
      description.foreach { __v =>
        _output__.writeString(6, __v)
      };
      elements.foreach { __v =>
        _output__.writeTag(8, 2)
        _output__.writeUInt32NoTag(__v.serializedSize)
        __v.writeTo(_output__)
      };
      title.foreach { __v =>
        _output__.writeString(9, __v)
      };
      color.foreach { __v =>
        _output__.writeString(10, __v)
      };
      pretext.foreach { __v =>
        _output__.writeString(11, __v)
      };
      `type`.foreach { __v =>
        _output__.writeString(12, __v)
      };
      size.foreach { __v =>
        _output__.writeInt32(13, __v)
      };
      length.foreach { __v =>
        _output__.writeInt32(15, __v)
      };
      buttons.foreach { __v =>
        _output__.writeTag(16, 2)
        _output__.writeUInt32NoTag(__v.serializedSize)
        __v.writeTo(_output__)
      };
      askInfo.foreach { __v =>
        _output__.writeTag(17, 2)
        _output__.writeUInt32NoTag(__v.serializedSize)
        __v.writeTo(_output__)
      };
      askInfoAnswer.foreach { __v =>
        _output__.writeTag(18, 2)
        _output__.writeUInt32NoTag(__v.serializedSize)
        __v.writeTo(_output__)
      };
      form.foreach { __v =>
        _output__.writeTag(20, 2)
        _output__.writeUInt32NoTag(__v.serializedSize)
        __v.writeTo(_output__)
      };
      formSubmit.foreach { __v =>
        _output__.writeTag(21, 2)
        _output__.writeUInt32NoTag(__v.serializedSize)
        __v.writeTo(_output__)
      };
    }
    def mergeFrom(`_input__`: _root_.com.google.protobuf.CodedInputStream): conversation.Attachment = {
      var __type = this.`type`
      var __mimetype = this.mimetype
      var __url = this.url
      var __thumbnailUrl = this.thumbnailUrl
      var __name = this.name
      var __description = this.description
      var __length = this.length
      var __size = this.size
      val __elements = (_root_.scala.collection.immutable.Vector.newBuilder[conversation.GenericElementTemplate] ++= this.elements)
      var __title = this.title
      var __color = this.color
      var __pretext = this.pretext
      val __buttons = (_root_.scala.collection.immutable.Vector.newBuilder[conversation.Button] ++= this.buttons)
      var __askInfo = this.askInfo
      var __askInfoAnswer = this.askInfoAnswer
      var __form = this.form
      var __formSubmit = this.formSubmit
      var _done__ = false
      while (!_done__) {
        val _tag__ = _input__.readTag()
        _tag__ match {
          case 0 => _done__ = true
          case 98 =>
            __type = Option(_input__.readString())
          case 18 =>
            __mimetype = Option(_input__.readString())
          case 26 =>
            __url = Option(_input__.readString())
          case 34 =>
            __thumbnailUrl = Option(_input__.readString())
          case 42 =>
            __name = Option(_input__.readString())
          case 50 =>
            __description = Option(_input__.readString())
          case 120 =>
            __length = Option(_input__.readInt32())
          case 104 =>
            __size = Option(_input__.readInt32())
          case 66 =>
            __elements += _root_.scalapb.LiteParser.readMessage(_input__, conversation.GenericElementTemplate.defaultInstance)
          case 74 =>
            __title = Option(_input__.readString())
          case 82 =>
            __color = Option(_input__.readString())
          case 90 =>
            __pretext = Option(_input__.readString())
          case 130 =>
            __buttons += _root_.scalapb.LiteParser.readMessage(_input__, conversation.Button.defaultInstance)
          case 138 =>
            __askInfo = Option(_root_.scalapb.LiteParser.readMessage(_input__, __askInfo.getOrElse(conversation.AskInfomation.defaultInstance)))
          case 146 =>
            __askInfoAnswer = Option(_root_.scalapb.LiteParser.readMessage(_input__, __askInfoAnswer.getOrElse(conversation.AskInfomationAnswer.defaultInstance)))
          case 162 =>
            __form = Option(_root_.scalapb.LiteParser.readMessage(_input__, __form.getOrElse(conversation.Form.defaultInstance)))
          case 170 =>
            __formSubmit = Option(_root_.scalapb.LiteParser.readMessage(_input__, __formSubmit.getOrElse(conversation.FormSubmit.defaultInstance)))
          case tag => _input__.skipField(tag)
        }
      }
      conversation.Attachment(
          `type` = __type,
          mimetype = __mimetype,
          url = __url,
          thumbnailUrl = __thumbnailUrl,
          name = __name,
          description = __description,
          length = __length,
          size = __size,
          elements = __elements.result(),
          title = __title,
          color = __color,
          pretext = __pretext,
          buttons = __buttons.result(),
          askInfo = __askInfo,
          askInfoAnswer = __askInfoAnswer,
          form = __form,
          formSubmit = __formSubmit
      )
    }
    def getType: _root_.scala.Predef.String = `type`.getOrElse("")
    def clearType: Attachment = copy(`type` = None)
    def withType(__v: _root_.scala.Predef.String): Attachment = copy(`type` = Option(__v))
    def getMimetype: _root_.scala.Predef.String = mimetype.getOrElse("")
    def clearMimetype: Attachment = copy(mimetype = None)
    def withMimetype(__v: _root_.scala.Predef.String): Attachment = copy(mimetype = Option(__v))
    def getUrl: _root_.scala.Predef.String = url.getOrElse("")
    def clearUrl: Attachment = copy(url = None)
    def withUrl(__v: _root_.scala.Predef.String): Attachment = copy(url = Option(__v))
    def getThumbnailUrl: _root_.scala.Predef.String = thumbnailUrl.getOrElse("")
    def clearThumbnailUrl: Attachment = copy(thumbnailUrl = None)
    def withThumbnailUrl(__v: _root_.scala.Predef.String): Attachment = copy(thumbnailUrl = Option(__v))
    def getName: _root_.scala.Predef.String = name.getOrElse("")
    def clearName: Attachment = copy(name = None)
    def withName(__v: _root_.scala.Predef.String): Attachment = copy(name = Option(__v))
    def getDescription: _root_.scala.Predef.String = description.getOrElse("")
    def clearDescription: Attachment = copy(description = None)
    def withDescription(__v: _root_.scala.Predef.String): Attachment = copy(description = Option(__v))
    def getLength: _root_.scala.Int = length.getOrElse(0)
    def clearLength: Attachment = copy(length = None)
    def withLength(__v: _root_.scala.Int): Attachment = copy(length = Option(__v))
    def getSize: _root_.scala.Int = size.getOrElse(0)
    def clearSize: Attachment = copy(size = None)
    def withSize(__v: _root_.scala.Int): Attachment = copy(size = Option(__v))
    def clearElements = copy(elements = _root_.scala.collection.Seq.empty)
    def addElements(__vs: conversation.GenericElementTemplate*): Attachment = addAllElements(__vs)
    def addAllElements(__vs: TraversableOnce[conversation.GenericElementTemplate]): Attachment = copy(elements = elements ++ __vs)
    def withElements(__v: _root_.scala.collection.Seq[conversation.GenericElementTemplate]): Attachment = copy(elements = __v)
    def getTitle: _root_.scala.Predef.String = title.getOrElse("")
    def clearTitle: Attachment = copy(title = None)
    def withTitle(__v: _root_.scala.Predef.String): Attachment = copy(title = Option(__v))
    def getColor: _root_.scala.Predef.String = color.getOrElse("")
    def clearColor: Attachment = copy(color = None)
    def withColor(__v: _root_.scala.Predef.String): Attachment = copy(color = Option(__v))
    def getPretext: _root_.scala.Predef.String = pretext.getOrElse("")
    def clearPretext: Attachment = copy(pretext = None)
    def withPretext(__v: _root_.scala.Predef.String): Attachment = copy(pretext = Option(__v))
    def clearButtons = copy(buttons = _root_.scala.collection.Seq.empty)
    def addButtons(__vs: conversation.Button*): Attachment = addAllButtons(__vs)
    def addAllButtons(__vs: TraversableOnce[conversation.Button]): Attachment = copy(buttons = buttons ++ __vs)
    def withButtons(__v: _root_.scala.collection.Seq[conversation.Button]): Attachment = copy(buttons = __v)
    def getAskInfo: conversation.AskInfomation = askInfo.getOrElse(conversation.AskInfomation.defaultInstance)
    def clearAskInfo: Attachment = copy(askInfo = None)
    def withAskInfo(__v: conversation.AskInfomation): Attachment = copy(askInfo = Option(__v))
    def getAskInfoAnswer: conversation.AskInfomationAnswer = askInfoAnswer.getOrElse(conversation.AskInfomationAnswer.defaultInstance)
    def clearAskInfoAnswer: Attachment = copy(askInfoAnswer = None)
    def withAskInfoAnswer(__v: conversation.AskInfomationAnswer): Attachment = copy(askInfoAnswer = Option(__v))
    def getForm: conversation.Form = form.getOrElse(conversation.Form.defaultInstance)
    def clearForm: Attachment = copy(form = None)
    def withForm(__v: conversation.Form): Attachment = copy(form = Option(__v))
    def getFormSubmit: conversation.FormSubmit = formSubmit.getOrElse(conversation.FormSubmit.defaultInstance)
    def clearFormSubmit: Attachment = copy(formSubmit = None)
    def withFormSubmit(__v: conversation.FormSubmit): Attachment = copy(formSubmit = Option(__v))
    def getFieldByNumber(__fieldNumber: _root_.scala.Int): scala.Any = {
      (__fieldNumber: @_root_.scala.unchecked) match {
        case 12 => `type`.orNull
        case 2 => mimetype.orNull
        case 3 => url.orNull
        case 4 => thumbnailUrl.orNull
        case 5 => name.orNull
        case 6 => description.orNull
        case 15 => length.orNull
        case 13 => size.orNull
        case 8 => elements
        case 9 => title.orNull
        case 10 => color.orNull
        case 11 => pretext.orNull
        case 16 => buttons
        case 17 => askInfo.orNull
        case 18 => askInfoAnswer.orNull
        case 20 => form.orNull
        case 21 => formSubmit.orNull
      }
    }
    def getField(__field: _root_.scalapb.descriptors.FieldDescriptor): _root_.scalapb.descriptors.PValue = {
      require(__field.containingMessage eq companion.scalaDescriptor)
      (__field.number: @_root_.scala.unchecked) match {
        case 12 => `type`.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 2 => mimetype.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 3 => url.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 4 => thumbnailUrl.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 5 => name.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 6 => description.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 15 => length.map(_root_.scalapb.descriptors.PInt).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 13 => size.map(_root_.scalapb.descriptors.PInt).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 8 => _root_.scalapb.descriptors.PRepeated(elements.map(_.toPMessage)(_root_.scala.collection.breakOut))
        case 9 => title.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 10 => color.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 11 => pretext.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 16 => _root_.scalapb.descriptors.PRepeated(buttons.map(_.toPMessage)(_root_.scala.collection.breakOut))
        case 17 => askInfo.map(_.toPMessage).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 18 => askInfoAnswer.map(_.toPMessage).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 20 => form.map(_.toPMessage).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 21 => formSubmit.map(_.toPMessage).getOrElse(_root_.scalapb.descriptors.PEmpty)
      }
    }
    def toProtoString: _root_.scala.Predef.String = _root_.scalapb.TextFormat.printToUnicodeString(this)
    def companion = conversation.Attachment
}

object Attachment extends scalapb.GeneratedMessageCompanion[conversation.Attachment] {
  implicit def messageCompanion: scalapb.GeneratedMessageCompanion[conversation.Attachment] = this
  def fromFieldsMap(__fieldsMap: scala.collection.immutable.Map[_root_.com.google.protobuf.Descriptors.FieldDescriptor, scala.Any]): conversation.Attachment = {
    require(__fieldsMap.keys.forall(_.getContainingType() == javaDescriptor), "FieldDescriptor does not match message type.")
    val __fields = javaDescriptor.getFields
    conversation.Attachment(
      __fieldsMap.get(__fields.get(0)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(1)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(2)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(3)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(4)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(5)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(6)).asInstanceOf[scala.Option[_root_.scala.Int]],
      __fieldsMap.get(__fields.get(7)).asInstanceOf[scala.Option[_root_.scala.Int]],
      __fieldsMap.getOrElse(__fields.get(8), Nil).asInstanceOf[_root_.scala.collection.Seq[conversation.GenericElementTemplate]],
      __fieldsMap.get(__fields.get(9)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(10)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(11)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.getOrElse(__fields.get(12), Nil).asInstanceOf[_root_.scala.collection.Seq[conversation.Button]],
      __fieldsMap.get(__fields.get(13)).asInstanceOf[scala.Option[conversation.AskInfomation]],
      __fieldsMap.get(__fields.get(14)).asInstanceOf[scala.Option[conversation.AskInfomationAnswer]],
      __fieldsMap.get(__fields.get(15)).asInstanceOf[scala.Option[conversation.Form]],
      __fieldsMap.get(__fields.get(16)).asInstanceOf[scala.Option[conversation.FormSubmit]]
    )
  }
  implicit def messageReads: _root_.scalapb.descriptors.Reads[conversation.Attachment] = _root_.scalapb.descriptors.Reads{
    case _root_.scalapb.descriptors.PMessage(__fieldsMap) =>
      require(__fieldsMap.keys.forall(_.containingMessage == scalaDescriptor), "FieldDescriptor does not match message type.")
      conversation.Attachment(
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(12).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(2).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(3).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(4).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(5).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(6).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(15).get).flatMap(_.as[scala.Option[_root_.scala.Int]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(13).get).flatMap(_.as[scala.Option[_root_.scala.Int]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(8).get).map(_.as[_root_.scala.collection.Seq[conversation.GenericElementTemplate]]).getOrElse(_root_.scala.collection.Seq.empty),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(9).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(10).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(11).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(16).get).map(_.as[_root_.scala.collection.Seq[conversation.Button]]).getOrElse(_root_.scala.collection.Seq.empty),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(17).get).flatMap(_.as[scala.Option[conversation.AskInfomation]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(18).get).flatMap(_.as[scala.Option[conversation.AskInfomationAnswer]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(20).get).flatMap(_.as[scala.Option[conversation.Form]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(21).get).flatMap(_.as[scala.Option[conversation.FormSubmit]])
      )
    case _ => throw new RuntimeException("Expected PMessage")
  }
  def javaDescriptor: _root_.com.google.protobuf.Descriptors.Descriptor = ConversationProto.javaDescriptor.getMessageTypes.get(30)
  def scalaDescriptor: _root_.scalapb.descriptors.Descriptor = ConversationProto.scalaDescriptor.messages(30)
  def messageCompanionForFieldNumber(__number: _root_.scala.Int): _root_.scalapb.GeneratedMessageCompanion[_] = {
    var __out: _root_.scalapb.GeneratedMessageCompanion[_] = null
    (__number: @_root_.scala.unchecked) match {
      case 8 => __out = conversation.GenericElementTemplate
      case 16 => __out = conversation.Button
      case 17 => __out = conversation.AskInfomation
      case 18 => __out = conversation.AskInfomationAnswer
      case 20 => __out = conversation.Form
      case 21 => __out = conversation.FormSubmit
    }
    __out
  }
  lazy val nestedMessagesCompanions: Seq[_root_.scalapb.GeneratedMessageCompanion[_]] = Seq.empty
  def enumCompanionForFieldNumber(__fieldNumber: _root_.scala.Int): _root_.scalapb.GeneratedEnumCompanion[_] = throw new MatchError(__fieldNumber)
  lazy val defaultInstance = conversation.Attachment(
  )
  implicit class AttachmentLens[UpperPB](_l: _root_.scalapb.lenses.Lens[UpperPB, conversation.Attachment]) extends _root_.scalapb.lenses.ObjectLens[UpperPB, conversation.Attachment](_l) {
    def `type`: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getType)((c_, f_) => c_.copy(`type` = Option(f_)))
    def optionalType: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.`type`)((c_, f_) => c_.copy(`type` = f_))
    def mimetype: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getMimetype)((c_, f_) => c_.copy(mimetype = Option(f_)))
    def optionalMimetype: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.mimetype)((c_, f_) => c_.copy(mimetype = f_))
    def url: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getUrl)((c_, f_) => c_.copy(url = Option(f_)))
    def optionalUrl: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.url)((c_, f_) => c_.copy(url = f_))
    def thumbnailUrl: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getThumbnailUrl)((c_, f_) => c_.copy(thumbnailUrl = Option(f_)))
    def optionalThumbnailUrl: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.thumbnailUrl)((c_, f_) => c_.copy(thumbnailUrl = f_))
    def name: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getName)((c_, f_) => c_.copy(name = Option(f_)))
    def optionalName: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.name)((c_, f_) => c_.copy(name = f_))
    def description: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getDescription)((c_, f_) => c_.copy(description = Option(f_)))
    def optionalDescription: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.description)((c_, f_) => c_.copy(description = f_))
    def length: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Int] = field(_.getLength)((c_, f_) => c_.copy(length = Option(f_)))
    def optionalLength: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Int]] = field(_.length)((c_, f_) => c_.copy(length = f_))
    def size: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Int] = field(_.getSize)((c_, f_) => c_.copy(size = Option(f_)))
    def optionalSize: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Int]] = field(_.size)((c_, f_) => c_.copy(size = f_))
    def elements: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.collection.Seq[conversation.GenericElementTemplate]] = field(_.elements)((c_, f_) => c_.copy(elements = f_))
    def title: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getTitle)((c_, f_) => c_.copy(title = Option(f_)))
    def optionalTitle: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.title)((c_, f_) => c_.copy(title = f_))
    def color: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getColor)((c_, f_) => c_.copy(color = Option(f_)))
    def optionalColor: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.color)((c_, f_) => c_.copy(color = f_))
    def pretext: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getPretext)((c_, f_) => c_.copy(pretext = Option(f_)))
    def optionalPretext: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.pretext)((c_, f_) => c_.copy(pretext = f_))
    def buttons: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.collection.Seq[conversation.Button]] = field(_.buttons)((c_, f_) => c_.copy(buttons = f_))
    def askInfo: _root_.scalapb.lenses.Lens[UpperPB, conversation.AskInfomation] = field(_.getAskInfo)((c_, f_) => c_.copy(askInfo = Option(f_)))
    def optionalAskInfo: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[conversation.AskInfomation]] = field(_.askInfo)((c_, f_) => c_.copy(askInfo = f_))
    def askInfoAnswer: _root_.scalapb.lenses.Lens[UpperPB, conversation.AskInfomationAnswer] = field(_.getAskInfoAnswer)((c_, f_) => c_.copy(askInfoAnswer = Option(f_)))
    def optionalAskInfoAnswer: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[conversation.AskInfomationAnswer]] = field(_.askInfoAnswer)((c_, f_) => c_.copy(askInfoAnswer = f_))
    def form: _root_.scalapb.lenses.Lens[UpperPB, conversation.Form] = field(_.getForm)((c_, f_) => c_.copy(form = Option(f_)))
    def optionalForm: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[conversation.Form]] = field(_.form)((c_, f_) => c_.copy(form = f_))
    def formSubmit: _root_.scalapb.lenses.Lens[UpperPB, conversation.FormSubmit] = field(_.getFormSubmit)((c_, f_) => c_.copy(formSubmit = Option(f_)))
    def optionalFormSubmit: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[conversation.FormSubmit]] = field(_.formSubmit)((c_, f_) => c_.copy(formSubmit = f_))
  }
  final val TYPE_FIELD_NUMBER = 12
  final val MIMETYPE_FIELD_NUMBER = 2
  final val URL_FIELD_NUMBER = 3
  final val THUMBNAIL_URL_FIELD_NUMBER = 4
  final val NAME_FIELD_NUMBER = 5
  final val DESCRIPTION_FIELD_NUMBER = 6
  final val LENGTH_FIELD_NUMBER = 15
  final val SIZE_FIELD_NUMBER = 13
  final val ELEMENTS_FIELD_NUMBER = 8
  final val TITLE_FIELD_NUMBER = 9
  final val COLOR_FIELD_NUMBER = 10
  final val PRETEXT_FIELD_NUMBER = 11
  final val BUTTONS_FIELD_NUMBER = 16
  final val ASK_INFO_FIELD_NUMBER = 17
  final val ASK_INFO_ANSWER_FIELD_NUMBER = 18
  final val FORM_FIELD_NUMBER = 20
  final val FORM_SUBMIT_FIELD_NUMBER = 21
}