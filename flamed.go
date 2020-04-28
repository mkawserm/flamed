package flamed

import internalStorage "github.com/mkawserm/flamed/pkg/storage"
import internalStoraged "github.com/mkawserm/flamed/pkg/storaged"

type Storage internalStorage.Storage
type Cluster internalStoraged.Cluster
type Storaged internalStoraged.Storaged
