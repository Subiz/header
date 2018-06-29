// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO2

package account

@SerialVersionUID(0L)
final case class Agents(
    ctx: scala.Option[common.Context] = None,
    agents: _root_.scala.collection.Seq[account.Agent] = _root_.scala.collection.Seq.empty
    ) extends scalapb.GeneratedMessage with scalapb.Message[Agents] with scalapb.lenses.Updatable[Agents] {
    @transient
    private[this] var __serializedSizeCachedValue: _root_.scala.Int = 0
    private[this] def __computeSerializedValue(): _root_.scala.Int = {
      var __size = 0
      if (ctx.isDefined) { __size += 1 + _root_.com.google.protobuf.CodedOutputStream.computeUInt32SizeNoTag(ctx.get.serializedSize) + ctx.get.serializedSize }
      agents.foreach(agents => __size += 1 + _root_.com.google.protobuf.CodedOutputStream.computeUInt32SizeNoTag(agents.serializedSize) + agents.serializedSize)
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
      agents.foreach { __v =>
        _output__.writeTag(2, 2)
        _output__.writeUInt32NoTag(__v.serializedSize)
        __v.writeTo(_output__)
      };
    }
    def mergeFrom(`_input__`: _root_.com.google.protobuf.CodedInputStream): account.Agents = {
      var __ctx = this.ctx
      val __agents = (_root_.scala.collection.immutable.Vector.newBuilder[account.Agent] ++= this.agents)
      var _done__ = false
      while (!_done__) {
        val _tag__ = _input__.readTag()
        _tag__ match {
          case 0 => _done__ = true
          case 10 =>
            __ctx = Option(_root_.scalapb.LiteParser.readMessage(_input__, __ctx.getOrElse(common.Context.defaultInstance)))
          case 18 =>
            __agents += _root_.scalapb.LiteParser.readMessage(_input__, account.Agent.defaultInstance)
          case tag => _input__.skipField(tag)
        }
      }
      account.Agents(
          ctx = __ctx,
          agents = __agents.result()
      )
    }
    def getCtx: common.Context = ctx.getOrElse(common.Context.defaultInstance)
    def clearCtx: Agents = copy(ctx = None)
    def withCtx(__v: common.Context): Agents = copy(ctx = Option(__v))
    def clearAgents = copy(agents = _root_.scala.collection.Seq.empty)
    def addAgents(__vs: account.Agent*): Agents = addAllAgents(__vs)
    def addAllAgents(__vs: TraversableOnce[account.Agent]): Agents = copy(agents = agents ++ __vs)
    def withAgents(__v: _root_.scala.collection.Seq[account.Agent]): Agents = copy(agents = __v)
    def getFieldByNumber(__fieldNumber: _root_.scala.Int): scala.Any = {
      (__fieldNumber: @_root_.scala.unchecked) match {
        case 1 => ctx.orNull
        case 2 => agents
      }
    }
    def getField(__field: _root_.scalapb.descriptors.FieldDescriptor): _root_.scalapb.descriptors.PValue = {
      require(__field.containingMessage eq companion.scalaDescriptor)
      (__field.number: @_root_.scala.unchecked) match {
        case 1 => ctx.map(_.toPMessage).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 2 => _root_.scalapb.descriptors.PRepeated(agents.map(_.toPMessage)(_root_.scala.collection.breakOut))
      }
    }
    def toProtoString: _root_.scala.Predef.String = _root_.scalapb.TextFormat.printToUnicodeString(this)
    def companion = account.Agents
}

object Agents extends scalapb.GeneratedMessageCompanion[account.Agents] {
  implicit def messageCompanion: scalapb.GeneratedMessageCompanion[account.Agents] = this
  def fromFieldsMap(__fieldsMap: scala.collection.immutable.Map[_root_.com.google.protobuf.Descriptors.FieldDescriptor, scala.Any]): account.Agents = {
    require(__fieldsMap.keys.forall(_.getContainingType() == javaDescriptor), "FieldDescriptor does not match message type.")
    val __fields = javaDescriptor.getFields
    account.Agents(
      __fieldsMap.get(__fields.get(0)).asInstanceOf[scala.Option[common.Context]],
      __fieldsMap.getOrElse(__fields.get(1), Nil).asInstanceOf[_root_.scala.collection.Seq[account.Agent]]
    )
  }
  implicit def messageReads: _root_.scalapb.descriptors.Reads[account.Agents] = _root_.scalapb.descriptors.Reads{
    case _root_.scalapb.descriptors.PMessage(__fieldsMap) =>
      require(__fieldsMap.keys.forall(_.containingMessage == scalaDescriptor), "FieldDescriptor does not match message type.")
      account.Agents(
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(1).get).flatMap(_.as[scala.Option[common.Context]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(2).get).map(_.as[_root_.scala.collection.Seq[account.Agent]]).getOrElse(_root_.scala.collection.Seq.empty)
      )
    case _ => throw new RuntimeException("Expected PMessage")
  }
  def javaDescriptor: _root_.com.google.protobuf.Descriptors.Descriptor = AccountProto.javaDescriptor.getMessageTypes.get(9)
  def scalaDescriptor: _root_.scalapb.descriptors.Descriptor = AccountProto.scalaDescriptor.messages(9)
  def messageCompanionForFieldNumber(__number: _root_.scala.Int): _root_.scalapb.GeneratedMessageCompanion[_] = {
    var __out: _root_.scalapb.GeneratedMessageCompanion[_] = null
    (__number: @_root_.scala.unchecked) match {
      case 1 => __out = common.Context
      case 2 => __out = account.Agent
    }
    __out
  }
  lazy val nestedMessagesCompanions: Seq[_root_.scalapb.GeneratedMessageCompanion[_]] = Seq.empty
  def enumCompanionForFieldNumber(__fieldNumber: _root_.scala.Int): _root_.scalapb.GeneratedEnumCompanion[_] = throw new MatchError(__fieldNumber)
  lazy val defaultInstance = account.Agents(
  )
  implicit class AgentsLens[UpperPB](_l: _root_.scalapb.lenses.Lens[UpperPB, account.Agents]) extends _root_.scalapb.lenses.ObjectLens[UpperPB, account.Agents](_l) {
    def ctx: _root_.scalapb.lenses.Lens[UpperPB, common.Context] = field(_.getCtx)((c_, f_) => c_.copy(ctx = Option(f_)))
    def optionalCtx: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[common.Context]] = field(_.ctx)((c_, f_) => c_.copy(ctx = f_))
    def agents: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.collection.Seq[account.Agent]] = field(_.agents)((c_, f_) => c_.copy(agents = f_))
  }
  final val CTX_FIELD_NUMBER = 1
  final val AGENTS_FIELD_NUMBER = 2
}