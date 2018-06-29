// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO2

package conversation

@SerialVersionUID(0L)
final case class StartRequest(
    ctx: scala.Option[common.Context] = None,
    id: scala.Option[_root_.scala.Predef.String] = None,
    accountId: scala.Option[_root_.scala.Predef.String] = None,
    channelType: scala.Option[_root_.scala.Predef.String] = None,
    from: scala.Option[_root_.scala.Predef.String] = None,
    to: _root_.scala.collection.Seq[_root_.scala.Predef.String] = _root_.scala.collection.Seq.empty,
    pageUrl: scala.Option[_root_.scala.Predef.String] = None,
    pageTitle: scala.Option[_root_.scala.Predef.String] = None,
    message: scala.Option[_root_.scala.Predef.String] = None,
    browserLanguage: scala.Option[_root_.scala.Predef.String] = None,
    language: scala.Option[_root_.scala.Predef.String] = None,
    deviceType: scala.Option[_root_.scala.Predef.String] = None,
    created: scala.Option[_root_.scala.Long] = None,
    conversationId: scala.Option[_root_.scala.Predef.String] = None,
    ip: scala.Option[_root_.scala.Predef.String] = None,
    country: scala.Option[_root_.scala.Predef.String] = None,
    countryCode: scala.Option[_root_.scala.Predef.String] = None,
    city: scala.Option[_root_.scala.Predef.String] = None,
    timezone: scala.Option[_root_.scala.Predef.String] = None,
    starterId: scala.Option[_root_.scala.Predef.String] = None,
    starterType: scala.Option[_root_.scala.Predef.String] = None,
    agentIds: _root_.scala.collection.Seq[_root_.scala.Predef.String] = _root_.scala.collection.Seq.empty,
    user: scala.Option[_root_.user.User] = None,
    integrationId: scala.Option[_root_.scala.Predef.String] = None
    ) extends scalapb.GeneratedMessage with scalapb.Message[StartRequest] with scalapb.lenses.Updatable[StartRequest] {
    @transient
    private[this] var __serializedSizeCachedValue: _root_.scala.Int = 0
    private[this] def __computeSerializedValue(): _root_.scala.Int = {
      var __size = 0
      if (ctx.isDefined) { __size += 1 + _root_.com.google.protobuf.CodedOutputStream.computeUInt32SizeNoTag(ctx.get.serializedSize) + ctx.get.serializedSize }
      if (id.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(2, id.get) }
      if (accountId.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(3, accountId.get) }
      if (channelType.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(4, channelType.get) }
      if (from.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(12, from.get) }
      to.foreach(to => __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(5, to))
      if (pageUrl.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(6, pageUrl.get) }
      if (pageTitle.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(7, pageTitle.get) }
      if (message.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(8, message.get) }
      if (browserLanguage.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(9, browserLanguage.get) }
      if (language.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(10, language.get) }
      if (deviceType.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(11, deviceType.get) }
      if (created.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeInt64Size(13, created.get) }
      if (conversationId.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(14, conversationId.get) }
      if (ip.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(15, ip.get) }
      if (country.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(16, country.get) }
      if (countryCode.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(17, countryCode.get) }
      if (city.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(18, city.get) }
      if (timezone.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(19, timezone.get) }
      if (starterId.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(20, starterId.get) }
      if (starterType.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(21, starterType.get) }
      agentIds.foreach(agentIds => __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(22, agentIds))
      if (user.isDefined) { __size += 2 + _root_.com.google.protobuf.CodedOutputStream.computeUInt32SizeNoTag(user.get.serializedSize) + user.get.serializedSize }
      if (integrationId.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(26, integrationId.get) }
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
      id.foreach { __v =>
        _output__.writeString(2, __v)
      };
      accountId.foreach { __v =>
        _output__.writeString(3, __v)
      };
      channelType.foreach { __v =>
        _output__.writeString(4, __v)
      };
      to.foreach { __v =>
        _output__.writeString(5, __v)
      };
      pageUrl.foreach { __v =>
        _output__.writeString(6, __v)
      };
      pageTitle.foreach { __v =>
        _output__.writeString(7, __v)
      };
      message.foreach { __v =>
        _output__.writeString(8, __v)
      };
      browserLanguage.foreach { __v =>
        _output__.writeString(9, __v)
      };
      language.foreach { __v =>
        _output__.writeString(10, __v)
      };
      deviceType.foreach { __v =>
        _output__.writeString(11, __v)
      };
      from.foreach { __v =>
        _output__.writeString(12, __v)
      };
      created.foreach { __v =>
        _output__.writeInt64(13, __v)
      };
      conversationId.foreach { __v =>
        _output__.writeString(14, __v)
      };
      ip.foreach { __v =>
        _output__.writeString(15, __v)
      };
      country.foreach { __v =>
        _output__.writeString(16, __v)
      };
      countryCode.foreach { __v =>
        _output__.writeString(17, __v)
      };
      city.foreach { __v =>
        _output__.writeString(18, __v)
      };
      timezone.foreach { __v =>
        _output__.writeString(19, __v)
      };
      starterId.foreach { __v =>
        _output__.writeString(20, __v)
      };
      starterType.foreach { __v =>
        _output__.writeString(21, __v)
      };
      agentIds.foreach { __v =>
        _output__.writeString(22, __v)
      };
      user.foreach { __v =>
        _output__.writeTag(25, 2)
        _output__.writeUInt32NoTag(__v.serializedSize)
        __v.writeTo(_output__)
      };
      integrationId.foreach { __v =>
        _output__.writeString(26, __v)
      };
    }
    def mergeFrom(`_input__`: _root_.com.google.protobuf.CodedInputStream): conversation.StartRequest = {
      var __ctx = this.ctx
      var __id = this.id
      var __accountId = this.accountId
      var __channelType = this.channelType
      var __from = this.from
      val __to = (_root_.scala.collection.immutable.Vector.newBuilder[_root_.scala.Predef.String] ++= this.to)
      var __pageUrl = this.pageUrl
      var __pageTitle = this.pageTitle
      var __message = this.message
      var __browserLanguage = this.browserLanguage
      var __language = this.language
      var __deviceType = this.deviceType
      var __created = this.created
      var __conversationId = this.conversationId
      var __ip = this.ip
      var __country = this.country
      var __countryCode = this.countryCode
      var __city = this.city
      var __timezone = this.timezone
      var __starterId = this.starterId
      var __starterType = this.starterType
      val __agentIds = (_root_.scala.collection.immutable.Vector.newBuilder[_root_.scala.Predef.String] ++= this.agentIds)
      var __user = this.user
      var __integrationId = this.integrationId
      var _done__ = false
      while (!_done__) {
        val _tag__ = _input__.readTag()
        _tag__ match {
          case 0 => _done__ = true
          case 10 =>
            __ctx = Option(_root_.scalapb.LiteParser.readMessage(_input__, __ctx.getOrElse(common.Context.defaultInstance)))
          case 18 =>
            __id = Option(_input__.readString())
          case 26 =>
            __accountId = Option(_input__.readString())
          case 34 =>
            __channelType = Option(_input__.readString())
          case 98 =>
            __from = Option(_input__.readString())
          case 42 =>
            __to += _input__.readString()
          case 50 =>
            __pageUrl = Option(_input__.readString())
          case 58 =>
            __pageTitle = Option(_input__.readString())
          case 66 =>
            __message = Option(_input__.readString())
          case 74 =>
            __browserLanguage = Option(_input__.readString())
          case 82 =>
            __language = Option(_input__.readString())
          case 90 =>
            __deviceType = Option(_input__.readString())
          case 104 =>
            __created = Option(_input__.readInt64())
          case 114 =>
            __conversationId = Option(_input__.readString())
          case 122 =>
            __ip = Option(_input__.readString())
          case 130 =>
            __country = Option(_input__.readString())
          case 138 =>
            __countryCode = Option(_input__.readString())
          case 146 =>
            __city = Option(_input__.readString())
          case 154 =>
            __timezone = Option(_input__.readString())
          case 162 =>
            __starterId = Option(_input__.readString())
          case 170 =>
            __starterType = Option(_input__.readString())
          case 178 =>
            __agentIds += _input__.readString()
          case 202 =>
            __user = Option(_root_.scalapb.LiteParser.readMessage(_input__, __user.getOrElse(_root_.user.User.defaultInstance)))
          case 210 =>
            __integrationId = Option(_input__.readString())
          case tag => _input__.skipField(tag)
        }
      }
      conversation.StartRequest(
          ctx = __ctx,
          id = __id,
          accountId = __accountId,
          channelType = __channelType,
          from = __from,
          to = __to.result(),
          pageUrl = __pageUrl,
          pageTitle = __pageTitle,
          message = __message,
          browserLanguage = __browserLanguage,
          language = __language,
          deviceType = __deviceType,
          created = __created,
          conversationId = __conversationId,
          ip = __ip,
          country = __country,
          countryCode = __countryCode,
          city = __city,
          timezone = __timezone,
          starterId = __starterId,
          starterType = __starterType,
          agentIds = __agentIds.result(),
          user = __user,
          integrationId = __integrationId
      )
    }
    def getCtx: common.Context = ctx.getOrElse(common.Context.defaultInstance)
    def clearCtx: StartRequest = copy(ctx = None)
    def withCtx(__v: common.Context): StartRequest = copy(ctx = Option(__v))
    def getId: _root_.scala.Predef.String = id.getOrElse("")
    def clearId: StartRequest = copy(id = None)
    def withId(__v: _root_.scala.Predef.String): StartRequest = copy(id = Option(__v))
    def getAccountId: _root_.scala.Predef.String = accountId.getOrElse("")
    def clearAccountId: StartRequest = copy(accountId = None)
    def withAccountId(__v: _root_.scala.Predef.String): StartRequest = copy(accountId = Option(__v))
    def getChannelType: _root_.scala.Predef.String = channelType.getOrElse("")
    def clearChannelType: StartRequest = copy(channelType = None)
    def withChannelType(__v: _root_.scala.Predef.String): StartRequest = copy(channelType = Option(__v))
    def getFrom: _root_.scala.Predef.String = from.getOrElse("")
    def clearFrom: StartRequest = copy(from = None)
    def withFrom(__v: _root_.scala.Predef.String): StartRequest = copy(from = Option(__v))
    def clearTo = copy(to = _root_.scala.collection.Seq.empty)
    def addTo(__vs: _root_.scala.Predef.String*): StartRequest = addAllTo(__vs)
    def addAllTo(__vs: TraversableOnce[_root_.scala.Predef.String]): StartRequest = copy(to = to ++ __vs)
    def withTo(__v: _root_.scala.collection.Seq[_root_.scala.Predef.String]): StartRequest = copy(to = __v)
    def getPageUrl: _root_.scala.Predef.String = pageUrl.getOrElse("")
    def clearPageUrl: StartRequest = copy(pageUrl = None)
    def withPageUrl(__v: _root_.scala.Predef.String): StartRequest = copy(pageUrl = Option(__v))
    def getPageTitle: _root_.scala.Predef.String = pageTitle.getOrElse("")
    def clearPageTitle: StartRequest = copy(pageTitle = None)
    def withPageTitle(__v: _root_.scala.Predef.String): StartRequest = copy(pageTitle = Option(__v))
    def getMessage: _root_.scala.Predef.String = message.getOrElse("")
    def clearMessage: StartRequest = copy(message = None)
    def withMessage(__v: _root_.scala.Predef.String): StartRequest = copy(message = Option(__v))
    def getBrowserLanguage: _root_.scala.Predef.String = browserLanguage.getOrElse("")
    def clearBrowserLanguage: StartRequest = copy(browserLanguage = None)
    def withBrowserLanguage(__v: _root_.scala.Predef.String): StartRequest = copy(browserLanguage = Option(__v))
    def getLanguage: _root_.scala.Predef.String = language.getOrElse("")
    def clearLanguage: StartRequest = copy(language = None)
    def withLanguage(__v: _root_.scala.Predef.String): StartRequest = copy(language = Option(__v))
    def getDeviceType: _root_.scala.Predef.String = deviceType.getOrElse("")
    def clearDeviceType: StartRequest = copy(deviceType = None)
    def withDeviceType(__v: _root_.scala.Predef.String): StartRequest = copy(deviceType = Option(__v))
    def getCreated: _root_.scala.Long = created.getOrElse(0L)
    def clearCreated: StartRequest = copy(created = None)
    def withCreated(__v: _root_.scala.Long): StartRequest = copy(created = Option(__v))
    def getConversationId: _root_.scala.Predef.String = conversationId.getOrElse("")
    def clearConversationId: StartRequest = copy(conversationId = None)
    def withConversationId(__v: _root_.scala.Predef.String): StartRequest = copy(conversationId = Option(__v))
    def getIp: _root_.scala.Predef.String = ip.getOrElse("")
    def clearIp: StartRequest = copy(ip = None)
    def withIp(__v: _root_.scala.Predef.String): StartRequest = copy(ip = Option(__v))
    def getCountry: _root_.scala.Predef.String = country.getOrElse("")
    def clearCountry: StartRequest = copy(country = None)
    def withCountry(__v: _root_.scala.Predef.String): StartRequest = copy(country = Option(__v))
    def getCountryCode: _root_.scala.Predef.String = countryCode.getOrElse("")
    def clearCountryCode: StartRequest = copy(countryCode = None)
    def withCountryCode(__v: _root_.scala.Predef.String): StartRequest = copy(countryCode = Option(__v))
    def getCity: _root_.scala.Predef.String = city.getOrElse("")
    def clearCity: StartRequest = copy(city = None)
    def withCity(__v: _root_.scala.Predef.String): StartRequest = copy(city = Option(__v))
    def getTimezone: _root_.scala.Predef.String = timezone.getOrElse("")
    def clearTimezone: StartRequest = copy(timezone = None)
    def withTimezone(__v: _root_.scala.Predef.String): StartRequest = copy(timezone = Option(__v))
    def getStarterId: _root_.scala.Predef.String = starterId.getOrElse("")
    def clearStarterId: StartRequest = copy(starterId = None)
    def withStarterId(__v: _root_.scala.Predef.String): StartRequest = copy(starterId = Option(__v))
    def getStarterType: _root_.scala.Predef.String = starterType.getOrElse("")
    def clearStarterType: StartRequest = copy(starterType = None)
    def withStarterType(__v: _root_.scala.Predef.String): StartRequest = copy(starterType = Option(__v))
    def clearAgentIds = copy(agentIds = _root_.scala.collection.Seq.empty)
    def addAgentIds(__vs: _root_.scala.Predef.String*): StartRequest = addAllAgentIds(__vs)
    def addAllAgentIds(__vs: TraversableOnce[_root_.scala.Predef.String]): StartRequest = copy(agentIds = agentIds ++ __vs)
    def withAgentIds(__v: _root_.scala.collection.Seq[_root_.scala.Predef.String]): StartRequest = copy(agentIds = __v)
    def getUser: _root_.user.User = user.getOrElse(_root_.user.User.defaultInstance)
    def clearUser: StartRequest = copy(user = None)
    def withUser(__v: _root_.user.User): StartRequest = copy(user = Option(__v))
    def getIntegrationId: _root_.scala.Predef.String = integrationId.getOrElse("")
    def clearIntegrationId: StartRequest = copy(integrationId = None)
    def withIntegrationId(__v: _root_.scala.Predef.String): StartRequest = copy(integrationId = Option(__v))
    def getFieldByNumber(__fieldNumber: _root_.scala.Int): scala.Any = {
      (__fieldNumber: @_root_.scala.unchecked) match {
        case 1 => ctx.orNull
        case 2 => id.orNull
        case 3 => accountId.orNull
        case 4 => channelType.orNull
        case 12 => from.orNull
        case 5 => to
        case 6 => pageUrl.orNull
        case 7 => pageTitle.orNull
        case 8 => message.orNull
        case 9 => browserLanguage.orNull
        case 10 => language.orNull
        case 11 => deviceType.orNull
        case 13 => created.orNull
        case 14 => conversationId.orNull
        case 15 => ip.orNull
        case 16 => country.orNull
        case 17 => countryCode.orNull
        case 18 => city.orNull
        case 19 => timezone.orNull
        case 20 => starterId.orNull
        case 21 => starterType.orNull
        case 22 => agentIds
        case 25 => user.orNull
        case 26 => integrationId.orNull
      }
    }
    def getField(__field: _root_.scalapb.descriptors.FieldDescriptor): _root_.scalapb.descriptors.PValue = {
      require(__field.containingMessage eq companion.scalaDescriptor)
      (__field.number: @_root_.scala.unchecked) match {
        case 1 => ctx.map(_.toPMessage).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 2 => id.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 3 => accountId.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 4 => channelType.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 12 => from.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 5 => _root_.scalapb.descriptors.PRepeated(to.map(_root_.scalapb.descriptors.PString)(_root_.scala.collection.breakOut))
        case 6 => pageUrl.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 7 => pageTitle.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 8 => message.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 9 => browserLanguage.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 10 => language.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 11 => deviceType.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 13 => created.map(_root_.scalapb.descriptors.PLong).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 14 => conversationId.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 15 => ip.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 16 => country.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 17 => countryCode.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 18 => city.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 19 => timezone.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 20 => starterId.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 21 => starterType.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 22 => _root_.scalapb.descriptors.PRepeated(agentIds.map(_root_.scalapb.descriptors.PString)(_root_.scala.collection.breakOut))
        case 25 => user.map(_.toPMessage).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 26 => integrationId.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
      }
    }
    def toProtoString: _root_.scala.Predef.String = _root_.scalapb.TextFormat.printToUnicodeString(this)
    def companion = conversation.StartRequest
}

object StartRequest extends scalapb.GeneratedMessageCompanion[conversation.StartRequest] {
  implicit def messageCompanion: scalapb.GeneratedMessageCompanion[conversation.StartRequest] = this
  def fromFieldsMap(__fieldsMap: scala.collection.immutable.Map[_root_.com.google.protobuf.Descriptors.FieldDescriptor, scala.Any]): conversation.StartRequest = {
    require(__fieldsMap.keys.forall(_.getContainingType() == javaDescriptor), "FieldDescriptor does not match message type.")
    val __fields = javaDescriptor.getFields
    conversation.StartRequest(
      __fieldsMap.get(__fields.get(0)).asInstanceOf[scala.Option[common.Context]],
      __fieldsMap.get(__fields.get(1)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(2)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(3)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(4)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.getOrElse(__fields.get(5), Nil).asInstanceOf[_root_.scala.collection.Seq[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(6)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(7)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(8)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(9)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(10)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(11)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(12)).asInstanceOf[scala.Option[_root_.scala.Long]],
      __fieldsMap.get(__fields.get(13)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(14)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(15)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(16)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(17)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(18)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(19)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(20)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.getOrElse(__fields.get(21), Nil).asInstanceOf[_root_.scala.collection.Seq[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(22)).asInstanceOf[scala.Option[_root_.user.User]],
      __fieldsMap.get(__fields.get(23)).asInstanceOf[scala.Option[_root_.scala.Predef.String]]
    )
  }
  implicit def messageReads: _root_.scalapb.descriptors.Reads[conversation.StartRequest] = _root_.scalapb.descriptors.Reads{
    case _root_.scalapb.descriptors.PMessage(__fieldsMap) =>
      require(__fieldsMap.keys.forall(_.containingMessage == scalaDescriptor), "FieldDescriptor does not match message type.")
      conversation.StartRequest(
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(1).get).flatMap(_.as[scala.Option[common.Context]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(2).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(3).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(4).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(12).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(5).get).map(_.as[_root_.scala.collection.Seq[_root_.scala.Predef.String]]).getOrElse(_root_.scala.collection.Seq.empty),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(6).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(7).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(8).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(9).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(10).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(11).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(13).get).flatMap(_.as[scala.Option[_root_.scala.Long]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(14).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(15).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(16).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(17).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(18).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(19).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(20).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(21).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(22).get).map(_.as[_root_.scala.collection.Seq[_root_.scala.Predef.String]]).getOrElse(_root_.scala.collection.Seq.empty),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(25).get).flatMap(_.as[scala.Option[_root_.user.User]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(26).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]])
      )
    case _ => throw new RuntimeException("Expected PMessage")
  }
  def javaDescriptor: _root_.com.google.protobuf.Descriptors.Descriptor = ConversationProto.javaDescriptor.getMessageTypes.get(11)
  def scalaDescriptor: _root_.scalapb.descriptors.Descriptor = ConversationProto.scalaDescriptor.messages(11)
  def messageCompanionForFieldNumber(__number: _root_.scala.Int): _root_.scalapb.GeneratedMessageCompanion[_] = {
    var __out: _root_.scalapb.GeneratedMessageCompanion[_] = null
    (__number: @_root_.scala.unchecked) match {
      case 1 => __out = common.Context
      case 25 => __out = user.User
    }
    __out
  }
  lazy val nestedMessagesCompanions: Seq[_root_.scalapb.GeneratedMessageCompanion[_]] = Seq.empty
  def enumCompanionForFieldNumber(__fieldNumber: _root_.scala.Int): _root_.scalapb.GeneratedEnumCompanion[_] = throw new MatchError(__fieldNumber)
  lazy val defaultInstance = conversation.StartRequest(
  )
  implicit class StartRequestLens[UpperPB](_l: _root_.scalapb.lenses.Lens[UpperPB, conversation.StartRequest]) extends _root_.scalapb.lenses.ObjectLens[UpperPB, conversation.StartRequest](_l) {
    def ctx: _root_.scalapb.lenses.Lens[UpperPB, common.Context] = field(_.getCtx)((c_, f_) => c_.copy(ctx = Option(f_)))
    def optionalCtx: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[common.Context]] = field(_.ctx)((c_, f_) => c_.copy(ctx = f_))
    def id: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getId)((c_, f_) => c_.copy(id = Option(f_)))
    def optionalId: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.id)((c_, f_) => c_.copy(id = f_))
    def accountId: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getAccountId)((c_, f_) => c_.copy(accountId = Option(f_)))
    def optionalAccountId: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.accountId)((c_, f_) => c_.copy(accountId = f_))
    def channelType: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getChannelType)((c_, f_) => c_.copy(channelType = Option(f_)))
    def optionalChannelType: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.channelType)((c_, f_) => c_.copy(channelType = f_))
    def from: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getFrom)((c_, f_) => c_.copy(from = Option(f_)))
    def optionalFrom: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.from)((c_, f_) => c_.copy(from = f_))
    def to: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.collection.Seq[_root_.scala.Predef.String]] = field(_.to)((c_, f_) => c_.copy(to = f_))
    def pageUrl: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getPageUrl)((c_, f_) => c_.copy(pageUrl = Option(f_)))
    def optionalPageUrl: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.pageUrl)((c_, f_) => c_.copy(pageUrl = f_))
    def pageTitle: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getPageTitle)((c_, f_) => c_.copy(pageTitle = Option(f_)))
    def optionalPageTitle: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.pageTitle)((c_, f_) => c_.copy(pageTitle = f_))
    def message: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getMessage)((c_, f_) => c_.copy(message = Option(f_)))
    def optionalMessage: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.message)((c_, f_) => c_.copy(message = f_))
    def browserLanguage: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getBrowserLanguage)((c_, f_) => c_.copy(browserLanguage = Option(f_)))
    def optionalBrowserLanguage: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.browserLanguage)((c_, f_) => c_.copy(browserLanguage = f_))
    def language: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getLanguage)((c_, f_) => c_.copy(language = Option(f_)))
    def optionalLanguage: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.language)((c_, f_) => c_.copy(language = f_))
    def deviceType: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getDeviceType)((c_, f_) => c_.copy(deviceType = Option(f_)))
    def optionalDeviceType: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.deviceType)((c_, f_) => c_.copy(deviceType = f_))
    def created: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Long] = field(_.getCreated)((c_, f_) => c_.copy(created = Option(f_)))
    def optionalCreated: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Long]] = field(_.created)((c_, f_) => c_.copy(created = f_))
    def conversationId: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getConversationId)((c_, f_) => c_.copy(conversationId = Option(f_)))
    def optionalConversationId: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.conversationId)((c_, f_) => c_.copy(conversationId = f_))
    def ip: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getIp)((c_, f_) => c_.copy(ip = Option(f_)))
    def optionalIp: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.ip)((c_, f_) => c_.copy(ip = f_))
    def country: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getCountry)((c_, f_) => c_.copy(country = Option(f_)))
    def optionalCountry: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.country)((c_, f_) => c_.copy(country = f_))
    def countryCode: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getCountryCode)((c_, f_) => c_.copy(countryCode = Option(f_)))
    def optionalCountryCode: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.countryCode)((c_, f_) => c_.copy(countryCode = f_))
    def city: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getCity)((c_, f_) => c_.copy(city = Option(f_)))
    def optionalCity: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.city)((c_, f_) => c_.copy(city = f_))
    def timezone: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getTimezone)((c_, f_) => c_.copy(timezone = Option(f_)))
    def optionalTimezone: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.timezone)((c_, f_) => c_.copy(timezone = f_))
    def starterId: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getStarterId)((c_, f_) => c_.copy(starterId = Option(f_)))
    def optionalStarterId: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.starterId)((c_, f_) => c_.copy(starterId = f_))
    def starterType: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getStarterType)((c_, f_) => c_.copy(starterType = Option(f_)))
    def optionalStarterType: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.starterType)((c_, f_) => c_.copy(starterType = f_))
    def agentIds: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.collection.Seq[_root_.scala.Predef.String]] = field(_.agentIds)((c_, f_) => c_.copy(agentIds = f_))
    def user: _root_.scalapb.lenses.Lens[UpperPB, _root_.user.User] = field(_.getUser)((c_, f_) => c_.copy(user = Option(f_)))
    def optionalUser: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.user.User]] = field(_.user)((c_, f_) => c_.copy(user = f_))
    def integrationId: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getIntegrationId)((c_, f_) => c_.copy(integrationId = Option(f_)))
    def optionalIntegrationId: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.integrationId)((c_, f_) => c_.copy(integrationId = f_))
  }
  final val CTX_FIELD_NUMBER = 1
  final val ID_FIELD_NUMBER = 2
  final val ACCOUNT_ID_FIELD_NUMBER = 3
  final val CHANNEL_TYPE_FIELD_NUMBER = 4
  final val FROM_FIELD_NUMBER = 12
  final val TO_FIELD_NUMBER = 5
  final val PAGE_URL_FIELD_NUMBER = 6
  final val PAGE_TITLE_FIELD_NUMBER = 7
  final val MESSAGE_FIELD_NUMBER = 8
  final val BROWSER_LANGUAGE_FIELD_NUMBER = 9
  final val LANGUAGE_FIELD_NUMBER = 10
  final val DEVICE_TYPE_FIELD_NUMBER = 11
  final val CREATED_FIELD_NUMBER = 13
  final val CONVERSATION_ID_FIELD_NUMBER = 14
  final val IP_FIELD_NUMBER = 15
  final val COUNTRY_FIELD_NUMBER = 16
  final val COUNTRY_CODE_FIELD_NUMBER = 17
  final val CITY_FIELD_NUMBER = 18
  final val TIMEZONE_FIELD_NUMBER = 19
  final val STARTER_ID_FIELD_NUMBER = 20
  final val STARTER_TYPE_FIELD_NUMBER = 21
  final val AGENT_IDS_FIELD_NUMBER = 22
  final val USER_FIELD_NUMBER = 25
  final val INTEGRATION_ID_FIELD_NUMBER = 26
}