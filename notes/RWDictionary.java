package dictionary;

class RWDictionary {
	// State variable:
	private final Map<String, String> m = new TreeMap<>();
	// Sync. variables:
	private final ReentrantReadWriteLock rwl =
		new ReentrantReadWriteLock();
	private final Lock r = rwl.readLock();
	private final Lock w = rwl.writeLock();

	public String get(String key) {
		r.lock();
		try {
			return m.get(key);
		} finally {
			r.unlock();
		}
	}

	public String put(String key, String val) {
		w.lock();
		try {
			return m.put(key, val);
		} finally {
			w.unlock();
		}
	}
}