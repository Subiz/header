// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO2

package conversation

/** @param title
  *   post back
  * @param url
  *  	optional RawEvent event = 12;
  */
@SerialVersionUID(0L)
final case class Button(
    `type`: scala.Option[_root_.scala.Predef.String] = None,
    id: scala.Option[_root_.scala.Predef.String] = None,
    title: scala.Option[_root_.scala.Predef.String] = None,
    payload: scala.Option[_root_.scala.Predef.String] = None,
    imageUrl: scala.Option[_root_.scala.Predef.String] = None,
    contentId: scala.Option[_root_.scala.Predef.String] = None,
    url: scala.Option[_root_.scala.Predef.String] = None
    ) extends scalapb.GeneratedMessage with scalapb.Message[Button] with scalapb.lenses.Updatable[Button] {
    @transient
    private[this] var __serializedSizeCachedValue: _root_.scala.Int = 0
    private[this] def __computeSerializedValue(): _root_.scala.Int = {
      var __size = 0
      if (`type`.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(2, `type`.get) }
      if (id.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(15, id.get) }
      if (title.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(3, title.get) }
      if (payload.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(4, payload.get) }
      if (imageUrl.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(5, imageUrl.get) }
      if (contentId.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(10, contentId.get) }
      if (url.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(14, url.get) }
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
      `type`.foreach { __v =>
        _output__.writeString(2, __v)
      };
      title.foreach { __v =>
        _output__.writeString(3, __v)
      };
      payload.foreach { __v =>
        _output__.writeString(4, __v)
      };
      imageUrl.foreach { __v =>
        _output__.writeString(5, __v)
      };
      contentId.foreach { __v =>
        _output__.writeString(10, __v)
      };
      url.foreach { __v =>
        _output__.writeString(14, __v)
      };
      id.foreach { __v =>
        _output__.writeString(15, __v)
      };
    }
    def mergeFrom(`_input__`: _root_.com.google.protobuf.CodedInputStream): conversation.Button = {
      var __type = this.`type`
      var __id = this.id
      var __title = this.title
      var __payload = this.payload
      var __imageUrl = this.imageUrl
      var __contentId = this.contentId
      var __url = this.url
      var _done__ = false
      while (!_done__) {
        val _tag__ = _input__.readTag()
        _tag__ match {
          case 0 => _done__ = true
          case 18 =>
            __type = Option(_input__.readString())
          case 122 =>
            __id = Option(_input__.readString())
          case 26 =>
            __title = Option(_input__.readString())
          case 34 =>
            __payload = Option(_input__.readString())
          case 42 =>
            __imageUrl = Option(_input__.readString())
          case 82 =>
            __contentId = Option(_input__.readString())
          case 114 =>
            __url = Option(_input__.readString())
          case tag => _input__.skipField(tag)
        }
      }
      conversation.Button(
          `type` = __type,
          id = __id,
          title = __title,
          payload = __payload,
          imageUrl = __imageUrl,
          contentId = __contentId,
          url = __url
      )
    }
    def getType: _root_.scala.Predef.String = `type`.getOrElse("")
    def clearType: Button = copy(`type` = None)
    def withType(__v: _root_.scala.Predef.String): Button = copy(`type` = Option(__v))
    def getId: _root_.scala.Predef.String = id.getOrElse("")
    def clearId: Button = copy(id = None)
    def withId(__v: _root_.scala.Predef.String): Button = copy(id = Option(__v))
    def getTitle: _root_.scala.Predef.String = title.getOrElse("")
    def clearTitle: Button = copy(title = None)
    def withTitle(__v: _root_.scala.Predef.String): Button = copy(title = Option(__v))
    def getPayload: _root_.scala.Predef.String = payload.getOrElse("")
    def clearPayload: Button = copy(payload = None)
    def withPayload(__v: _root_.scala.Predef.String): Button = copy(payload = Option(__v))
    def getImageUrl: _root_.scala.Predef.String = imageUrl.getOrElse("")
    def clearImageUrl: Button = copy(imageUrl = None)
    def withImageUrl(__v: _root_.scala.Predef.String): Button = copy(imageUrl = Option(__v))
    def getContentId: _root_.scala.Predef.String = contentId.getOrElse("")
    def clearContentId: Button = copy(contentId = None)
    def withContentId(__v: _root_.scala.Predef.String): Button = copy(contentId = Option(__v))
    def getUrl: _root_.scala.Predef.String = url.getOrElse("")
    def clearUrl: Button = copy(url = None)
    def withUrl(__v: _root_.scala.Predef.String): Button = copy(url = Option(__v))
    def getFieldByNumber(__fieldNumber: _root_.scala.Int): scala.Any = {
      (__fieldNumber: @_root_.scala.unchecked) match {
        case 2 => `type`.orNull
        case 15 => id.orNull
        case 3 => title.orNull
        case 4 => payload.orNull
        case 5 => imageUrl.orNull
        case 10 => contentId.orNull
        case 14 => url.orNull
      }
    }
    def getField(__field: _root_.scalapb.descriptors.FieldDescriptor): _root_.scalapb.descriptors.PValue = {
      require(__field.containingMessage eq companion.scalaDescriptor)
      (__field.number: @_root_.scala.unchecked) match {
        case 2 => `type`.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 15 => id.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 3 => title.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 4 => payload.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 5 => imageUrl.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 10 => contentId.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 14 => url.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
      }
    }
    def toProtoString: _root_.scala.Predef.String = _root_.scalapb.TextFormat.printToUnicodeString(this)
    def companion = conversation.Button
}

object Button extends scalapb.GeneratedMessageCompanion[conversation.Button] {
  implicit def messageCompanion: scalapb.GeneratedMessageCompanion[conversation.Button] = this
  def fromFieldsMap(__fieldsMap: scala.collection.immutable.Map[_root_.com.google.protobuf.Descriptors.FieldDescriptor, scala.Any]): conversation.Button = {
    require(__fieldsMap.keys.forall(_.getContainingType() == javaDescriptor), "FieldDescriptor does not match message type.")
    val __fields = javaDescriptor.getFields
    conversation.Button(
      __fieldsMap.get(__fields.get(0)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(1)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(2)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(3)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(4)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(5)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(6)).asInstanceOf[scala.Option[_root_.scala.Predef.String]]
    )
  }
  implicit def messageReads: _root_.scalapb.descriptors.Reads[conversation.Button] = _root_.scalapb.descriptors.Reads{
    case _root_.scalapb.descriptors.PMessage(__fieldsMap) =>
      require(__fieldsMap.keys.forall(_.containingMessage == scalaDescriptor), "FieldDescriptor does not match message type.")
      conversation.Button(
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(2).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(15).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(3).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(4).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(5).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(10).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(14).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]])
      )
    case _ => throw new RuntimeException("Expected PMessage")
  }
  def javaDescriptor: _root_.com.google.protobuf.Descriptors.Descriptor = ConversationProto.javaDescriptor.getMessageTypes.get(26)
  def scalaDescriptor: _root_.scalapb.descriptors.Descriptor = ConversationProto.scalaDescriptor.messages(26)
  def messageCompanionForFieldNumber(__number: _root_.scala.Int): _root_.scalapb.GeneratedMessageCompanion[_] = throw new MatchError(__number)
  lazy val nestedMessagesCompanions: Seq[_root_.scalapb.GeneratedMessageCompanion[_]] = Seq.empty
  def enumCompanionForFieldNumber(__fieldNumber: _root_.scala.Int): _root_.scalapb.GeneratedEnumCompanion[_] = throw new MatchError(__fieldNumber)
  lazy val defaultInstance = conversation.Button(
  )
  sealed trait ButtonType extends _root_.scalapb.GeneratedEnum {
    type EnumType = ButtonType
    def isurlbutton: _root_.scala.Boolean = false
    def ispostbackbutton: _root_.scala.Boolean = false
    def iseventbutton: _root_.scala.Boolean = false
    def companion: _root_.scalapb.GeneratedEnumCompanion[ButtonType] = conversation.Button.ButtonType
  }
  
  object ButtonType extends _root_.scalapb.GeneratedEnumCompanion[ButtonType] {
    implicit def enumCompanion: _root_.scalapb.GeneratedEnumCompanion[ButtonType] = this
    @SerialVersionUID(0L)
    case object url_button extends ButtonType {
      val value = 2
      val index = 0
      val name = "url_button"
      override def isurlbutton: _root_.scala.Boolean = true
    }
    
    @SerialVersionUID(0L)
    case object postback_button extends ButtonType {
      val value = 3
      val index = 1
      val name = "postback_button"
      override def ispostbackbutton: _root_.scala.Boolean = true
    }
    
    @SerialVersionUID(0L)
    case object event_button extends ButtonType {
      val value = 4
      val index = 2
      val name = "event_button"
      override def iseventbutton: _root_.scala.Boolean = true
    }
    
    @SerialVersionUID(0L)
    final case class Unrecognized(value: _root_.scala.Int) extends ButtonType with _root_.scalapb.UnrecognizedEnum
    
    lazy val values = scala.collection.Seq(url_button, postback_button, event_button)
    def fromValue(value: _root_.scala.Int): ButtonType = value match {
      case 2 => url_button
      case 3 => postback_button
      case 4 => event_button
      case __other => Unrecognized(__other)
    }
    def javaDescriptor: _root_.com.google.protobuf.Descriptors.EnumDescriptor = conversation.Button.javaDescriptor.getEnumTypes.get(0)
    def scalaDescriptor: _root_.scalapb.descriptors.EnumDescriptor = conversation.Button.scalaDescriptor.enums(0)
  }
  implicit class ButtonLens[UpperPB](_l: _root_.scalapb.lenses.Lens[UpperPB, conversation.Button]) extends _root_.scalapb.lenses.ObjectLens[UpperPB, conversation.Button](_l) {
    def `type`: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getType)((c_, f_) => c_.copy(`type` = Option(f_)))
    def optionalType: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.`type`)((c_, f_) => c_.copy(`type` = f_))
    def id: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getId)((c_, f_) => c_.copy(id = Option(f_)))
    def optionalId: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.id)((c_, f_) => c_.copy(id = f_))
    def title: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getTitle)((c_, f_) => c_.copy(title = Option(f_)))
    def optionalTitle: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.title)((c_, f_) => c_.copy(title = f_))
    def payload: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getPayload)((c_, f_) => c_.copy(payload = Option(f_)))
    def optionalPayload: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.payload)((c_, f_) => c_.copy(payload = f_))
    def imageUrl: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getImageUrl)((c_, f_) => c_.copy(imageUrl = Option(f_)))
    def optionalImageUrl: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.imageUrl)((c_, f_) => c_.copy(imageUrl = f_))
    def contentId: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getContentId)((c_, f_) => c_.copy(contentId = Option(f_)))
    def optionalContentId: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.contentId)((c_, f_) => c_.copy(contentId = f_))
    def url: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getUrl)((c_, f_) => c_.copy(url = Option(f_)))
    def optionalUrl: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.url)((c_, f_) => c_.copy(url = f_))
  }
  final val TYPE_FIELD_NUMBER = 2
  final val ID_FIELD_NUMBER = 15
  final val TITLE_FIELD_NUMBER = 3
  final val PAYLOAD_FIELD_NUMBER = 4
  final val IMAGE_URL_FIELD_NUMBER = 5
  final val CONTENT_ID_FIELD_NUMBER = 10
  final val URL_FIELD_NUMBER = 14
}
