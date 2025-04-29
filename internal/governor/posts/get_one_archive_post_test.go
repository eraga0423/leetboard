package posts_governor

import "1337b0rd/internal/types/database"

type mockOnePostArchive struct {
	onePost database.ArchiveOnePostResp
}

func (m *mockOnePostArchive) OneArchivePost(database.ArchiveOnePostReq) (database.ArchiveOnePostResp, error) {
	return nil, nil
}
