// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO2

package mailkon

@SerialVersionUID(0L)
final case class Address(
    accountId: scala.Option[_root_.scala.Predef.String] = None,
    address: scala.Option[_root_.scala.Predef.String] = None,
    updated: scala.Option[_root_.scala.Long] = None
    ) extends scalapb.GeneratedMessage with scalapb.Message[Address] with scalapb.lenses.Updatable[Address] {
    @transient
    private[this] var __serializedSizeCachedValue: _root_.scala.Int = 0
    private[this] def __computeSerializedValue(): _root_.scala.Int = {
      var __size = 0
      if (accountId.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(1, accountId.get) }
      if (address.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(2, address.get) }
      if (updated.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeInt64Size(3, updated.get) }
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
      address.foreach { __v =>
        _output__.writeString(2, __v)
      };
      updated.foreach { __v =>
        _output__.writeInt64(3, __v)
      };
    }
    def mergeFrom(`_input__`: _root_.com.google.protobuf.CodedInputStream): mailkon.Address = {
      var __accountId = this.accountId
      var __address = this.address
      var __updated = this.updated
      var _done__ = false
      while (!_done__) {
        val _tag__ = _input__.readTag()
        _tag__ match {
          case 0 => _done__ = true
          case 10 =>
            __accountId = Option(_input__.readString())
          case 18 =>
            __address = Option(_input__.readString())
          case 24 =>
            __updated = Option(_input__.readInt64())
          case tag => _input__.skipField(tag)
        }
      }
      mailkon.Address(
          accountId = __accountId,
          address = __address,
          updated = __updated
      )
    }
    def getAccountId: _root_.scala.Predef.String = accountId.getOrElse("")
    def clearAccountId: Address = copy(accountId = None)
    def withAccountId(__v: _root_.scala.Predef.String): Address = copy(accountId = Option(__v))
    def getAddress: _root_.scala.Predef.String = address.getOrElse("")
    def clearAddress: Address = copy(address = None)
    def withAddress(__v: _root_.scala.Predef.String): Address = copy(address = Option(__v))
    def getUpdated: _root_.scala.Long = updated.getOrElse(0L)
    def clearUpdated: Address = copy(updated = None)
    def withUpdated(__v: _root_.scala.Long): Address = copy(updated = Option(__v))
    def getFieldByNumber(__fieldNumber: _root_.scala.Int): scala.Any = {
      (__fieldNumber: @_root_.scala.unchecked) match {
        case 1 => accountId.orNull
        case 2 => address.orNull
        case 3 => updated.orNull
      }
    }
    def getField(__field: _root_.scalapb.descriptors.FieldDescriptor): _root_.scalapb.descriptors.PValue = {
      require(__field.containingMessage eq companion.scalaDescriptor)
      (__field.number: @_root_.scala.unchecked) match {
        case 1 => accountId.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 2 => address.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 3 => updated.map(_root_.scalapb.descriptors.PLong).getOrElse(_root_.scalapb.descriptors.PEmpty)
      }
    }
    def toProtoString: _root_.scala.Predef.String = _root_.scalapb.TextFormat.printToUnicodeString(this)
    def companion = mailkon.Address
}

object Address extends scalapb.GeneratedMessageCompanion[mailkon.Address] {
  implicit def messageCompanion: scalapb.GeneratedMessageCompanion[mailkon.Address] = this
  def fromFieldsMap(__fieldsMap: scala.collection.immutable.Map[_root_.com.google.protobuf.Descriptors.FieldDescriptor, scala.Any]): mailkon.Address = {
    require(__fieldsMap.keys.forall(_.getContainingType() == javaDescriptor), "FieldDescriptor does not match message type.")
    val __fields = javaDescriptor.getFields
    mailkon.Address(
      __fieldsMap.get(__fields.get(0)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(1)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(2)).asInstanceOf[scala.Option[_root_.scala.Long]]
    )
  }
  implicit def messageReads: _root_.scalapb.descriptors.Reads[mailkon.Address] = _root_.scalapb.descriptors.Reads{
    case _root_.scalapb.descriptors.PMessage(__fieldsMap) =>
      require(__fieldsMap.keys.forall(_.containingMessage == scalaDescriptor), "FieldDescriptor does not match message type.")
      mailkon.Address(
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(1).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(2).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(3).get).flatMap(_.as[scala.Option[_root_.scala.Long]])
      )
    case _ => throw new RuntimeException("Expected PMessage")
  }
  def javaDescriptor: _root_.com.google.protobuf.Descriptors.Descriptor = MailkonProto.javaDescriptor.getMessageTypes.get(0)
  def scalaDescriptor: _root_.scalapb.descriptors.Descriptor = MailkonProto.scalaDescriptor.messages(0)
  def messageCompanionForFieldNumber(__number: _root_.scala.Int): _root_.scalapb.GeneratedMessageCompanion[_] = throw new MatchError(__number)
  lazy val nestedMessagesCompanions: Seq[_root_.scalapb.GeneratedMessageCompanion[_]] = Seq.empty
  def enumCompanionForFieldNumber(__fieldNumber: _root_.scala.Int): _root_.scalapb.GeneratedEnumCompanion[_] = throw new MatchError(__fieldNumber)
  lazy val defaultInstance = mailkon.Address(
  )
  implicit class AddressLens[UpperPB](_l: _root_.scalapb.lenses.Lens[UpperPB, mailkon.Address]) extends _root_.scalapb.lenses.ObjectLens[UpperPB, mailkon.Address](_l) {
    def accountId: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getAccountId)((c_, f_) => c_.copy(accountId = Option(f_)))
    def optionalAccountId: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.accountId)((c_, f_) => c_.copy(accountId = f_))
    def address: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getAddress)((c_, f_) => c_.copy(address = Option(f_)))
    def optionalAddress: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.address)((c_, f_) => c_.copy(address = f_))
    def updated: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Long] = field(_.getUpdated)((c_, f_) => c_.copy(updated = Option(f_)))
    def optionalUpdated: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Long]] = field(_.updated)((c_, f_) => c_.copy(updated = f_))
  }
  final val ACCOUNT_ID_FIELD_NUMBER = 1
  final val ADDRESS_FIELD_NUMBER = 2
  final val UPDATED_FIELD_NUMBER = 3
}