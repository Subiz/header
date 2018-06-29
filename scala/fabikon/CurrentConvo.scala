// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO2

package fabikon

@SerialVersionUID(0L)
final case class CurrentConvo(
    accountId: scala.Option[_root_.scala.Predef.String] = None,
    convoId: scala.Option[_root_.scala.Predef.String] = None,
    pageId: scala.Option[_root_.scala.Predef.String] = None,
    fbuserId: scala.Option[_root_.scala.Predef.String] = None,
    sbuserId: scala.Option[_root_.scala.Predef.String] = None
    ) extends scalapb.GeneratedMessage with scalapb.Message[CurrentConvo] with scalapb.lenses.Updatable[CurrentConvo] {
    @transient
    private[this] var __serializedSizeCachedValue: _root_.scala.Int = 0
    private[this] def __computeSerializedValue(): _root_.scala.Int = {
      var __size = 0
      if (accountId.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(2, accountId.get) }
      if (convoId.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(3, convoId.get) }
      if (pageId.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(4, pageId.get) }
      if (fbuserId.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(5, fbuserId.get) }
      if (sbuserId.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(6, sbuserId.get) }
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
        _output__.writeString(2, __v)
      };
      convoId.foreach { __v =>
        _output__.writeString(3, __v)
      };
      pageId.foreach { __v =>
        _output__.writeString(4, __v)
      };
      fbuserId.foreach { __v =>
        _output__.writeString(5, __v)
      };
      sbuserId.foreach { __v =>
        _output__.writeString(6, __v)
      };
    }
    def mergeFrom(`_input__`: _root_.com.google.protobuf.CodedInputStream): fabikon.CurrentConvo = {
      var __accountId = this.accountId
      var __convoId = this.convoId
      var __pageId = this.pageId
      var __fbuserId = this.fbuserId
      var __sbuserId = this.sbuserId
      var _done__ = false
      while (!_done__) {
        val _tag__ = _input__.readTag()
        _tag__ match {
          case 0 => _done__ = true
          case 18 =>
            __accountId = Option(_input__.readString())
          case 26 =>
            __convoId = Option(_input__.readString())
          case 34 =>
            __pageId = Option(_input__.readString())
          case 42 =>
            __fbuserId = Option(_input__.readString())
          case 50 =>
            __sbuserId = Option(_input__.readString())
          case tag => _input__.skipField(tag)
        }
      }
      fabikon.CurrentConvo(
          accountId = __accountId,
          convoId = __convoId,
          pageId = __pageId,
          fbuserId = __fbuserId,
          sbuserId = __sbuserId
      )
    }
    def getAccountId: _root_.scala.Predef.String = accountId.getOrElse("")
    def clearAccountId: CurrentConvo = copy(accountId = None)
    def withAccountId(__v: _root_.scala.Predef.String): CurrentConvo = copy(accountId = Option(__v))
    def getConvoId: _root_.scala.Predef.String = convoId.getOrElse("")
    def clearConvoId: CurrentConvo = copy(convoId = None)
    def withConvoId(__v: _root_.scala.Predef.String): CurrentConvo = copy(convoId = Option(__v))
    def getPageId: _root_.scala.Predef.String = pageId.getOrElse("")
    def clearPageId: CurrentConvo = copy(pageId = None)
    def withPageId(__v: _root_.scala.Predef.String): CurrentConvo = copy(pageId = Option(__v))
    def getFbuserId: _root_.scala.Predef.String = fbuserId.getOrElse("")
    def clearFbuserId: CurrentConvo = copy(fbuserId = None)
    def withFbuserId(__v: _root_.scala.Predef.String): CurrentConvo = copy(fbuserId = Option(__v))
    def getSbuserId: _root_.scala.Predef.String = sbuserId.getOrElse("")
    def clearSbuserId: CurrentConvo = copy(sbuserId = None)
    def withSbuserId(__v: _root_.scala.Predef.String): CurrentConvo = copy(sbuserId = Option(__v))
    def getFieldByNumber(__fieldNumber: _root_.scala.Int): scala.Any = {
      (__fieldNumber: @_root_.scala.unchecked) match {
        case 2 => accountId.orNull
        case 3 => convoId.orNull
        case 4 => pageId.orNull
        case 5 => fbuserId.orNull
        case 6 => sbuserId.orNull
      }
    }
    def getField(__field: _root_.scalapb.descriptors.FieldDescriptor): _root_.scalapb.descriptors.PValue = {
      require(__field.containingMessage eq companion.scalaDescriptor)
      (__field.number: @_root_.scala.unchecked) match {
        case 2 => accountId.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 3 => convoId.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 4 => pageId.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 5 => fbuserId.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 6 => sbuserId.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
      }
    }
    def toProtoString: _root_.scala.Predef.String = _root_.scalapb.TextFormat.printToUnicodeString(this)
    def companion = fabikon.CurrentConvo
}

object CurrentConvo extends scalapb.GeneratedMessageCompanion[fabikon.CurrentConvo] {
  implicit def messageCompanion: scalapb.GeneratedMessageCompanion[fabikon.CurrentConvo] = this
  def fromFieldsMap(__fieldsMap: scala.collection.immutable.Map[_root_.com.google.protobuf.Descriptors.FieldDescriptor, scala.Any]): fabikon.CurrentConvo = {
    require(__fieldsMap.keys.forall(_.getContainingType() == javaDescriptor), "FieldDescriptor does not match message type.")
    val __fields = javaDescriptor.getFields
    fabikon.CurrentConvo(
      __fieldsMap.get(__fields.get(0)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(1)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(2)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(3)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(4)).asInstanceOf[scala.Option[_root_.scala.Predef.String]]
    )
  }
  implicit def messageReads: _root_.scalapb.descriptors.Reads[fabikon.CurrentConvo] = _root_.scalapb.descriptors.Reads{
    case _root_.scalapb.descriptors.PMessage(__fieldsMap) =>
      require(__fieldsMap.keys.forall(_.containingMessage == scalaDescriptor), "FieldDescriptor does not match message type.")
      fabikon.CurrentConvo(
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(2).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(3).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(4).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(5).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(6).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]])
      )
    case _ => throw new RuntimeException("Expected PMessage")
  }
  def javaDescriptor: _root_.com.google.protobuf.Descriptors.Descriptor = FabikonProto.javaDescriptor.getMessageTypes.get(15)
  def scalaDescriptor: _root_.scalapb.descriptors.Descriptor = FabikonProto.scalaDescriptor.messages(15)
  def messageCompanionForFieldNumber(__number: _root_.scala.Int): _root_.scalapb.GeneratedMessageCompanion[_] = throw new MatchError(__number)
  lazy val nestedMessagesCompanions: Seq[_root_.scalapb.GeneratedMessageCompanion[_]] = Seq.empty
  def enumCompanionForFieldNumber(__fieldNumber: _root_.scala.Int): _root_.scalapb.GeneratedEnumCompanion[_] = throw new MatchError(__fieldNumber)
  lazy val defaultInstance = fabikon.CurrentConvo(
  )
  implicit class CurrentConvoLens[UpperPB](_l: _root_.scalapb.lenses.Lens[UpperPB, fabikon.CurrentConvo]) extends _root_.scalapb.lenses.ObjectLens[UpperPB, fabikon.CurrentConvo](_l) {
    def accountId: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getAccountId)((c_, f_) => c_.copy(accountId = Option(f_)))
    def optionalAccountId: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.accountId)((c_, f_) => c_.copy(accountId = f_))
    def convoId: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getConvoId)((c_, f_) => c_.copy(convoId = Option(f_)))
    def optionalConvoId: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.convoId)((c_, f_) => c_.copy(convoId = f_))
    def pageId: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getPageId)((c_, f_) => c_.copy(pageId = Option(f_)))
    def optionalPageId: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.pageId)((c_, f_) => c_.copy(pageId = f_))
    def fbuserId: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getFbuserId)((c_, f_) => c_.copy(fbuserId = Option(f_)))
    def optionalFbuserId: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.fbuserId)((c_, f_) => c_.copy(fbuserId = f_))
    def sbuserId: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getSbuserId)((c_, f_) => c_.copy(sbuserId = Option(f_)))
    def optionalSbuserId: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.sbuserId)((c_, f_) => c_.copy(sbuserId = f_))
  }
  final val ACCOUNT_ID_FIELD_NUMBER = 2
  final val CONVO_ID_FIELD_NUMBER = 3
  final val PAGE_ID_FIELD_NUMBER = 4
  final val FBUSER_ID_FIELD_NUMBER = 5
  final val SBUSER_ID_FIELD_NUMBER = 6
}