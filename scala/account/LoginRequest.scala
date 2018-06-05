// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO2

package account

@SerialVersionUID(0L)
final case class LoginRequest(
    ctx: scala.Option[common.Context] = None,
    email: scala.Option[_root_.scala.Predef.String] = None,
    password: scala.Option[_root_.scala.Predef.String] = None
    ) extends scalapb.GeneratedMessage with scalapb.Message[LoginRequest] with scalapb.lenses.Updatable[LoginRequest] {
    @transient
    private[this] var __serializedSizeCachedValue: _root_.scala.Int = 0
    private[this] def __computeSerializedValue(): _root_.scala.Int = {
      var __size = 0
      if (ctx.isDefined) { __size += 1 + _root_.com.google.protobuf.CodedOutputStream.computeUInt32SizeNoTag(ctx.get.serializedSize) + ctx.get.serializedSize }
      if (email.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(2, email.get) }
      if (password.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(3, password.get) }
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
      email.foreach { __v =>
        _output__.writeString(2, __v)
      };
      password.foreach { __v =>
        _output__.writeString(3, __v)
      };
    }
    def mergeFrom(`_input__`: _root_.com.google.protobuf.CodedInputStream): account.LoginRequest = {
      var __ctx = this.ctx
      var __email = this.email
      var __password = this.password
      var _done__ = false
      while (!_done__) {
        val _tag__ = _input__.readTag()
        _tag__ match {
          case 0 => _done__ = true
          case 10 =>
            __ctx = Option(_root_.scalapb.LiteParser.readMessage(_input__, __ctx.getOrElse(common.Context.defaultInstance)))
          case 18 =>
            __email = Option(_input__.readString())
          case 26 =>
            __password = Option(_input__.readString())
          case tag => _input__.skipField(tag)
        }
      }
      account.LoginRequest(
          ctx = __ctx,
          email = __email,
          password = __password
      )
    }
    def getCtx: common.Context = ctx.getOrElse(common.Context.defaultInstance)
    def clearCtx: LoginRequest = copy(ctx = None)
    def withCtx(__v: common.Context): LoginRequest = copy(ctx = Option(__v))
    def getEmail: _root_.scala.Predef.String = email.getOrElse("")
    def clearEmail: LoginRequest = copy(email = None)
    def withEmail(__v: _root_.scala.Predef.String): LoginRequest = copy(email = Option(__v))
    def getPassword: _root_.scala.Predef.String = password.getOrElse("")
    def clearPassword: LoginRequest = copy(password = None)
    def withPassword(__v: _root_.scala.Predef.String): LoginRequest = copy(password = Option(__v))
    def getFieldByNumber(__fieldNumber: _root_.scala.Int): scala.Any = {
      (__fieldNumber: @_root_.scala.unchecked) match {
        case 1 => ctx.orNull
        case 2 => email.orNull
        case 3 => password.orNull
      }
    }
    def getField(__field: _root_.scalapb.descriptors.FieldDescriptor): _root_.scalapb.descriptors.PValue = {
      require(__field.containingMessage eq companion.scalaDescriptor)
      (__field.number: @_root_.scala.unchecked) match {
        case 1 => ctx.map(_.toPMessage).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 2 => email.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 3 => password.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
      }
    }
    def toProtoString: _root_.scala.Predef.String = _root_.scalapb.TextFormat.printToUnicodeString(this)
    def companion = account.LoginRequest
}

object LoginRequest extends scalapb.GeneratedMessageCompanion[account.LoginRequest] {
  implicit def messageCompanion: scalapb.GeneratedMessageCompanion[account.LoginRequest] = this
  def fromFieldsMap(__fieldsMap: scala.collection.immutable.Map[_root_.com.google.protobuf.Descriptors.FieldDescriptor, scala.Any]): account.LoginRequest = {
    require(__fieldsMap.keys.forall(_.getContainingType() == javaDescriptor), "FieldDescriptor does not match message type.")
    val __fields = javaDescriptor.getFields
    account.LoginRequest(
      __fieldsMap.get(__fields.get(0)).asInstanceOf[scala.Option[common.Context]],
      __fieldsMap.get(__fields.get(1)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(2)).asInstanceOf[scala.Option[_root_.scala.Predef.String]]
    )
  }
  implicit def messageReads: _root_.scalapb.descriptors.Reads[account.LoginRequest] = _root_.scalapb.descriptors.Reads{
    case _root_.scalapb.descriptors.PMessage(__fieldsMap) =>
      require(__fieldsMap.keys.forall(_.containingMessage == scalaDescriptor), "FieldDescriptor does not match message type.")
      account.LoginRequest(
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(1).get).flatMap(_.as[scala.Option[common.Context]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(2).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(3).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]])
      )
    case _ => throw new RuntimeException("Expected PMessage")
  }
  def javaDescriptor: _root_.com.google.protobuf.Descriptors.Descriptor = AccountProto.javaDescriptor.getMessageTypes.get(8)
  def scalaDescriptor: _root_.scalapb.descriptors.Descriptor = AccountProto.scalaDescriptor.messages(8)
  def messageCompanionForFieldNumber(__number: _root_.scala.Int): _root_.scalapb.GeneratedMessageCompanion[_] = {
    var __out: _root_.scalapb.GeneratedMessageCompanion[_] = null
    (__number: @_root_.scala.unchecked) match {
      case 1 => __out = common.Context
    }
    __out
  }
  lazy val nestedMessagesCompanions: Seq[_root_.scalapb.GeneratedMessageCompanion[_]] = Seq.empty
  def enumCompanionForFieldNumber(__fieldNumber: _root_.scala.Int): _root_.scalapb.GeneratedEnumCompanion[_] = throw new MatchError(__fieldNumber)
  lazy val defaultInstance = account.LoginRequest(
  )
  implicit class LoginRequestLens[UpperPB](_l: _root_.scalapb.lenses.Lens[UpperPB, account.LoginRequest]) extends _root_.scalapb.lenses.ObjectLens[UpperPB, account.LoginRequest](_l) {
    def ctx: _root_.scalapb.lenses.Lens[UpperPB, common.Context] = field(_.getCtx)((c_, f_) => c_.copy(ctx = Option(f_)))
    def optionalCtx: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[common.Context]] = field(_.ctx)((c_, f_) => c_.copy(ctx = f_))
    def email: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getEmail)((c_, f_) => c_.copy(email = Option(f_)))
    def optionalEmail: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.email)((c_, f_) => c_.copy(email = f_))
    def password: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getPassword)((c_, f_) => c_.copy(password = Option(f_)))
    def optionalPassword: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.password)((c_, f_) => c_.copy(password = f_))
  }
  final val CTX_FIELD_NUMBER = 1
  final val EMAIL_FIELD_NUMBER = 2
  final val PASSWORD_FIELD_NUMBER = 3
}
