// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO2

package fabikon

@SerialVersionUID(0L)
final case class FacebookPage(
    accountId: scala.Option[_root_.scala.Predef.String] = None,
    id: scala.Option[_root_.scala.Predef.String] = None,
    created: scala.Option[_root_.scala.Long] = None,
    pictureUrl: scala.Option[_root_.scala.Predef.String] = None,
    name: scala.Option[_root_.scala.Predef.String] = None,
    accessToken: scala.Option[_root_.scala.Predef.String] = None
    ) extends scalapb.GeneratedMessage with scalapb.Message[FacebookPage] with scalapb.lenses.Updatable[FacebookPage] {
    @transient
    private[this] var __serializedSizeCachedValue: _root_.scala.Int = 0
    private[this] def __computeSerializedValue(): _root_.scala.Int = {
      var __size = 0
      if (accountId.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(1, accountId.get) }
      if (id.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(2, id.get) }
      if (created.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeInt64Size(3, created.get) }
      if (pictureUrl.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(4, pictureUrl.get) }
      if (name.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(5, name.get) }
      if (accessToken.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(6, accessToken.get) }
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
      accountId.foreach { __v =>
        _output__.writeString(1, __v)
      };
      id.foreach { __v =>
        _output__.writeString(2, __v)
      };
      created.foreach { __v =>
        _output__.writeInt64(3, __v)
      };
      pictureUrl.foreach { __v =>
        _output__.writeString(4, __v)
      };
      name.foreach { __v =>
        _output__.writeString(5, __v)
      };
      accessToken.foreach { __v =>
        _output__.writeString(6, __v)
      };
    }
    def mergeFrom(`_input__`: _root_.com.google.protobuf.CodedInputStream): fabikon.FacebookPage = {
      var __accountId = this.accountId
      var __id = this.id
      var __created = this.created
      var __pictureUrl = this.pictureUrl
      var __name = this.name
      var __accessToken = this.accessToken
      var _done__ = false
      while (!_done__) {
        val _tag__ = _input__.readTag()
        _tag__ match {
          case 0 => _done__ = true
          case 10 =>
            __accountId = Option(_input__.readString())
          case 18 =>
            __id = Option(_input__.readString())
          case 24 =>
            __created = Option(_input__.readInt64())
          case 34 =>
            __pictureUrl = Option(_input__.readString())
          case 42 =>
            __name = Option(_input__.readString())
          case 50 =>
            __accessToken = Option(_input__.readString())
          case tag => _input__.skipField(tag)
        }
      }
      fabikon.FacebookPage(
          accountId = __accountId,
          id = __id,
          created = __created,
          pictureUrl = __pictureUrl,
          name = __name,
          accessToken = __accessToken
      )
    }
    def getAccountId: _root_.scala.Predef.String = accountId.getOrElse("")
    def clearAccountId: FacebookPage = copy(accountId = None)
    def withAccountId(__v: _root_.scala.Predef.String): FacebookPage = copy(accountId = Option(__v))
    def getId: _root_.scala.Predef.String = id.getOrElse("")
    def clearId: FacebookPage = copy(id = None)
    def withId(__v: _root_.scala.Predef.String): FacebookPage = copy(id = Option(__v))
    def getCreated: _root_.scala.Long = created.getOrElse(0L)
    def clearCreated: FacebookPage = copy(created = None)
    def withCreated(__v: _root_.scala.Long): FacebookPage = copy(created = Option(__v))
    def getPictureUrl: _root_.scala.Predef.String = pictureUrl.getOrElse("")
    def clearPictureUrl: FacebookPage = copy(pictureUrl = None)
    def withPictureUrl(__v: _root_.scala.Predef.String): FacebookPage = copy(pictureUrl = Option(__v))
    def getName: _root_.scala.Predef.String = name.getOrElse("")
    def clearName: FacebookPage = copy(name = None)
    def withName(__v: _root_.scala.Predef.String): FacebookPage = copy(name = Option(__v))
    def getAccessToken: _root_.scala.Predef.String = accessToken.getOrElse("")
    def clearAccessToken: FacebookPage = copy(accessToken = None)
    def withAccessToken(__v: _root_.scala.Predef.String): FacebookPage = copy(accessToken = Option(__v))
    def getFieldByNumber(__fieldNumber: _root_.scala.Int): scala.Any = {
      (__fieldNumber: @_root_.scala.unchecked) match {
        case 1 => accountId.orNull
        case 2 => id.orNull
        case 3 => created.orNull
        case 4 => pictureUrl.orNull
        case 5 => name.orNull
        case 6 => accessToken.orNull
      }
    }
    def getField(__field: _root_.scalapb.descriptors.FieldDescriptor): _root_.scalapb.descriptors.PValue = {
      require(__field.containingMessage eq companion.scalaDescriptor)
      (__field.number: @_root_.scala.unchecked) match {
        case 1 => accountId.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 2 => id.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 3 => created.map(_root_.scalapb.descriptors.PLong).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 4 => pictureUrl.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 5 => name.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 6 => accessToken.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
      }
    }
    def toProtoString: _root_.scala.Predef.String = _root_.scalapb.TextFormat.printToUnicodeString(this)
    def companion = fabikon.FacebookPage
}

object FacebookPage extends scalapb.GeneratedMessageCompanion[fabikon.FacebookPage] {
  implicit def messageCompanion: scalapb.GeneratedMessageCompanion[fabikon.FacebookPage] = this
  def fromFieldsMap(__fieldsMap: scala.collection.immutable.Map[_root_.com.google.protobuf.Descriptors.FieldDescriptor, scala.Any]): fabikon.FacebookPage = {
    require(__fieldsMap.keys.forall(_.getContainingType() == javaDescriptor), "FieldDescriptor does not match message type.")
    val __fields = javaDescriptor.getFields
    fabikon.FacebookPage(
      __fieldsMap.get(__fields.get(0)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(1)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(2)).asInstanceOf[scala.Option[_root_.scala.Long]],
      __fieldsMap.get(__fields.get(3)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(4)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(5)).asInstanceOf[scala.Option[_root_.scala.Predef.String]]
    )
  }
  implicit def messageReads: _root_.scalapb.descriptors.Reads[fabikon.FacebookPage] = _root_.scalapb.descriptors.Reads{
    case _root_.scalapb.descriptors.PMessage(__fieldsMap) =>
      require(__fieldsMap.keys.forall(_.containingMessage == scalaDescriptor), "FieldDescriptor does not match message type.")
      fabikon.FacebookPage(
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(1).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(2).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(3).get).flatMap(_.as[scala.Option[_root_.scala.Long]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(4).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(5).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(6).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]])
      )
    case _ => throw new RuntimeException("Expected PMessage")
  }
  def javaDescriptor: _root_.com.google.protobuf.Descriptors.Descriptor = FabikonProto.javaDescriptor.getMessageTypes.get(0)
  def scalaDescriptor: _root_.scalapb.descriptors.Descriptor = FabikonProto.scalaDescriptor.messages(0)
  def messageCompanionForFieldNumber(__number: _root_.scala.Int): _root_.scalapb.GeneratedMessageCompanion[_] = throw new MatchError(__number)
  lazy val nestedMessagesCompanions: Seq[_root_.scalapb.GeneratedMessageCompanion[_]] = Seq.empty
  def enumCompanionForFieldNumber(__fieldNumber: _root_.scala.Int): _root_.scalapb.GeneratedEnumCompanion[_] = throw new MatchError(__fieldNumber)
  lazy val defaultInstance = fabikon.FacebookPage(
  )
  implicit class FacebookPageLens[UpperPB](_l: _root_.scalapb.lenses.Lens[UpperPB, fabikon.FacebookPage]) extends _root_.scalapb.lenses.ObjectLens[UpperPB, fabikon.FacebookPage](_l) {
    def accountId: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getAccountId)((c_, f_) => c_.copy(accountId = Option(f_)))
    def optionalAccountId: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.accountId)((c_, f_) => c_.copy(accountId = f_))
    def id: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getId)((c_, f_) => c_.copy(id = Option(f_)))
    def optionalId: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.id)((c_, f_) => c_.copy(id = f_))
    def created: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Long] = field(_.getCreated)((c_, f_) => c_.copy(created = Option(f_)))
    def optionalCreated: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Long]] = field(_.created)((c_, f_) => c_.copy(created = f_))
    def pictureUrl: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getPictureUrl)((c_, f_) => c_.copy(pictureUrl = Option(f_)))
    def optionalPictureUrl: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.pictureUrl)((c_, f_) => c_.copy(pictureUrl = f_))
    def name: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getName)((c_, f_) => c_.copy(name = Option(f_)))
    def optionalName: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.name)((c_, f_) => c_.copy(name = f_))
    def accessToken: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getAccessToken)((c_, f_) => c_.copy(accessToken = Option(f_)))
    def optionalAccessToken: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.accessToken)((c_, f_) => c_.copy(accessToken = f_))
  }
  final val ACCOUNT_ID_FIELD_NUMBER = 1
  final val ID_FIELD_NUMBER = 2
  final val CREATED_FIELD_NUMBER = 3
  final val PICTURE_URL_FIELD_NUMBER = 4
  final val NAME_FIELD_NUMBER = 5
  final val ACCESS_TOKEN_FIELD_NUMBER = 6
}