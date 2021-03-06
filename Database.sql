USE [data7322]
GO
/****** Object:  User [bps7322]    Script Date: 12/18/2017 2:43:18 PM ******/
CREATE USER [bps7322] FOR LOGIN [bps7322] WITH DEFAULT_SCHEMA=[db_accessadmin]
GO
ALTER ROLE [db_accessadmin] ADD MEMBER [bps7322]
GO
ALTER ROLE [db_datareader] ADD MEMBER [bps7322]
GO
ALTER ROLE [db_datawriter] ADD MEMBER [bps7322]
GO
/****** Object:  Table [dbo].[Jabatan]    Script Date: 12/18/2017 2:43:19 PM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
SET ANSI_PADDING ON
GO
CREATE TABLE [dbo].[Jabatan](
	[ID] [int] IDENTITY(1,1) NOT NULL,
	[Deskripsi] [varchar](50) NULL,
 CONSTRAINT [PK_Jabatan] PRIMARY KEY CLUSTERED 
(
	[ID] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]

GO
SET ANSI_PADDING OFF
GO
/****** Object:  Table [dbo].[Kegiatan]    Script Date: 12/18/2017 2:43:19 PM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
SET ANSI_PADDING ON
GO
CREATE TABLE [dbo].[Kegiatan](
	[ID] [int] IDENTITY(1,1) NOT NULL,
	[Output] [int] NOT NULL,
	[Deskripsi] [varchar](50) NULL,
 CONSTRAINT [PK_Kegiatan] PRIMARY KEY CLUSTERED 
(
	[ID] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]

GO
SET ANSI_PADDING OFF
GO
/****** Object:  Table [dbo].[Output]    Script Date: 12/18/2017 2:43:19 PM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
SET ANSI_PADDING ON
GO
CREATE TABLE [dbo].[Output](
	[ID] [int] IDENTITY(1,1) NOT NULL,
	[Deskripsi] [varchar](50) NULL,
 CONSTRAINT [PK_Output] PRIMARY KEY CLUSTERED 
(
	[ID] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]

GO
SET ANSI_PADDING OFF
GO
/****** Object:  Table [dbo].[Pegawai]    Script Date: 12/18/2017 2:43:19 PM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
SET ANSI_PADDING ON
GO
CREATE TABLE [dbo].[Pegawai](
	[NIP] [varchar](18) NOT NULL,
	[Nama] [varchar](50) NOT NULL,
	[Jabatan] [int] NOT NULL,
	[Seksi] [int] NOT NULL,
	[Telepon] [varchar](20) NULL,
	[Email] [varchar](50) NULL,
 CONSTRAINT [PK_Pegawai] PRIMARY KEY CLUSTERED 
(
	[NIP] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]

GO
SET ANSI_PADDING OFF
GO
/****** Object:  Table [dbo].[Seksi]    Script Date: 12/18/2017 2:43:19 PM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
SET ANSI_PADDING ON
GO
CREATE TABLE [dbo].[Seksi](
	[ID] [int] IDENTITY(1,1) NOT NULL,
	[Deskripsi] [varchar](50) NULL,
 CONSTRAINT [PK_Seksi] PRIMARY KEY CLUSTERED 
(
	[ID] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]

GO
SET ANSI_PADDING OFF
GO
/****** Object:  Table [dbo].[Task]    Script Date: 12/18/2017 2:43:19 PM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
SET ANSI_PADDING ON
GO
CREATE TABLE [dbo].[Task](
	[ID] [int] IDENTITY(1,1) NOT NULL,
	[Kegiatan] [int] NOT NULL,
	[Jumlah] [int] NULL,
	[Mulai] [date] NULL,
	[Selesai] [date] NULL,
	[Seksi] [int] NOT NULL,
	[Deskripsi] [varchar](50) NULL,
 CONSTRAINT [PK_Task] PRIMARY KEY CLUSTERED 
(
	[ID] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]

GO
SET ANSI_PADDING OFF
GO
/****** Object:  Table [dbo].[TaskDesc]    Script Date: 12/18/2017 2:43:19 PM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
SET ANSI_PADDING ON
GO
CREATE TABLE [dbo].[TaskDesc](
	[ID] [int] IDENTITY(1,1) NOT NULL,
	[TaskID] [int] NOT NULL,
	[Author] [int] NOT NULL,
	[Judul] [varchar](50) NULL,
	[Deskripsi] [varchar](50) NULL,
	[Jumlah] [int] NULL,
	[isDelegated] [int] NULL,
 CONSTRAINT [PK_TaskDesc] PRIMARY KEY CLUSTERED 
(
	[ID] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]

GO
SET ANSI_PADDING OFF
GO
/****** Object:  Table [dbo].[TaskList]    Script Date: 12/18/2017 2:43:19 PM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
SET ANSI_PADDING ON
GO
CREATE TABLE [dbo].[TaskList](
	[ID] [int] IDENTITY(1,1) NOT NULL,
	[TaskID] [int] NOT NULL,
	[Pegawai] [varchar](18) NOT NULL,
	[Parent] [int] NOT NULL,
	[Mulai] [date] NULL,
	[Selesai] [date] NULL,
	[Status] [int] NOT NULL,
 CONSTRAINT [PK_TaskList] PRIMARY KEY CLUSTERED 
(
	[ID] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]

GO
SET ANSI_PADDING OFF
GO
ALTER TABLE [dbo].[Pegawai]  WITH CHECK ADD  CONSTRAINT [FK_Pegawai_Jabatan] FOREIGN KEY([Jabatan])
REFERENCES [dbo].[Jabatan] ([ID])
GO
ALTER TABLE [dbo].[Pegawai] CHECK CONSTRAINT [FK_Pegawai_Jabatan]
GO
ALTER TABLE [dbo].[Task]  WITH CHECK ADD  CONSTRAINT [FK_Task_Kegiatan] FOREIGN KEY([Kegiatan])
REFERENCES [dbo].[Kegiatan] ([ID])
GO
ALTER TABLE [dbo].[Task] CHECK CONSTRAINT [FK_Task_Kegiatan]
GO
ALTER TABLE [dbo].[Task]  WITH CHECK ADD  CONSTRAINT [FK_Task_Seksi] FOREIGN KEY([Seksi])
REFERENCES [dbo].[Seksi] ([ID])
GO
ALTER TABLE [dbo].[Task] CHECK CONSTRAINT [FK_Task_Seksi]
GO
ALTER TABLE [dbo].[TaskDesc]  WITH CHECK ADD  CONSTRAINT [FK_TaskDesc_Task] FOREIGN KEY([TaskID])
REFERENCES [dbo].[Task] ([ID])
GO
ALTER TABLE [dbo].[TaskDesc] CHECK CONSTRAINT [FK_TaskDesc_Task]
GO
ALTER TABLE [dbo].[TaskList]  WITH CHECK ADD  CONSTRAINT [FK_TaskList_Pegawai] FOREIGN KEY([Pegawai])
REFERENCES [dbo].[Pegawai] ([NIP])
GO
ALTER TABLE [dbo].[TaskList] CHECK CONSTRAINT [FK_TaskList_Pegawai]
GO
ALTER TABLE [dbo].[TaskList]  WITH CHECK ADD  CONSTRAINT [FK_TaskList_Seksi] FOREIGN KEY([Parent])
REFERENCES [dbo].[Seksi] ([ID])
GO
ALTER TABLE [dbo].[TaskList] CHECK CONSTRAINT [FK_TaskList_Seksi]
GO
ALTER TABLE [dbo].[TaskList]  WITH CHECK ADD  CONSTRAINT [FK_TaskList_TaskDesc] FOREIGN KEY([TaskID])
REFERENCES [dbo].[TaskDesc] ([ID])
GO
ALTER TABLE [dbo].[TaskList] CHECK CONSTRAINT [FK_TaskList_TaskDesc]
GO
EXEC sys.sp_addextendedproperty @name=N'MS_Description', @value=N'Menyimpan ID Jabatan' , @level0type=N'SCHEMA',@level0name=N'dbo', @level1type=N'TABLE',@level1name=N'Jabatan', @level2type=N'COLUMN',@level2name=N'ID'
GO
EXEC sys.sp_addextendedproperty @name=N'MS_Description', @value=N'FK ke Output kolom ID' , @level0type=N'SCHEMA',@level0name=N'dbo', @level1type=N'TABLE',@level1name=N'Kegiatan', @level2type=N'COLUMN',@level2name=N'Output'
GO
EXEC sys.sp_addextendedproperty @name=N'MS_Description', @value=N'Menyimpan ID Seksi' , @level0type=N'SCHEMA',@level0name=N'dbo', @level1type=N'TABLE',@level1name=N'Seksi', @level2type=N'COLUMN',@level2name=N'ID'
GO
EXEC sys.sp_addextendedproperty @name=N'MS_Description', @value=N'Menyimpan ID Tugas' , @level0type=N'SCHEMA',@level0name=N'dbo', @level1type=N'TABLE',@level1name=N'Task', @level2type=N'COLUMN',@level2name=N'ID'
GO
