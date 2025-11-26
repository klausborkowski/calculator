import { useState, useEffect } from 'react'
import './index.css'

// Use relative path for nginx proxy in Docker
// In development mode through Vite use empty string for proxy
// In production through nginx also use relative paths
const API_BASE_URL = import.meta.env.DEV ? '' : ''

function App() {
  const [packages, setPackages] = useState({})
  const [packageSizeInput, setPackageSizeInput] = useState('')
  const [orderSize, setOrderSize] = useState('')
  const [results, setResults] = useState({})
  const [swaggerOpen, setSwaggerOpen] = useState(false)

  // Load package list on component mount
  useEffect(() => {
    loadPackages()
  }, [])

  const loadPackages = async () => {
    try {
      const response = await fetch(`${API_BASE_URL}/packages`)
      if (response.ok) {
        const data = await response.json()
        setPackages(data)
      }
    } catch (error) {
      console.error('Error loading packages:', error)
    }
  }

  const addPackage = async () => {
    const value = parseInt(packageSizeInput)
    if (!value || value <= 0) {
      alert('Please enter a valid package size')
      return
    }

    try {
      const response = await fetch(`${API_BASE_URL}/package`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ packageSize: value }),
      })

      if (response.ok) {
        setPackageSizeInput('')
        loadPackages() // Reload package list
      } else {
        const errorText = await response.text()
        alert('Error adding package: ' + (errorText || response.status))
      }
    } catch (error) {
      alert('Error: ' + error.message)
    }
  }

  const removePackage = async (packageId) => {
    if (!packageId) {
      alert('Error deleting package: missing id')
      return
    }

    try {
      const response = await fetch(`${API_BASE_URL}/package/${packageId}`, {
        method: 'DELETE',
      })

      if (response.ok) {
        loadPackages() // Reload package list
      } else {
        const errorText = await response.text()
        alert('Error deleting package: ' + (errorText || response.status))
      }
    } catch (error) {
      alert('Error: ' + error.message)
    }
  }

  const calculate = async () => {
    const orderSizeValue = parseInt(orderSize)
    if (!orderSizeValue || orderSizeValue <= 0) {
      alert('Please enter a valid order size')
      return
    }

    try {
      const response = await fetch(`${API_BASE_URL}/calculate`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(orderSizeValue),
      })

      if (response.ok) {
        const result = await response.json()
        setResults(result)
      } else {
        alert('Error calculating result')
      }
    } catch (error) {
      alert('Error: ' + error.message)
    }
  }

  const openSwagger = () => {
    setSwaggerOpen(true)
  }

  const closeSwagger = () => {
    setSwaggerOpen(false)
  }

  return (
    <>
      <button className="swagger-btn" onClick={openSwagger}>
        Swagger
      </button>

      {/* Block 1: Packages */}
      <div className="container packages">
        <h3>Packages</h3>
        <div id="package-list">
          {Object.entries(packages).map(([id, size]) => (
            <div key={id} className="package-item">
              <input type="number" value={size} disabled />
              <button onClick={() => removePackage(id)}>❌</button>
            </div>
          ))}
        </div>
        <div id="new-package">
          <input
            type="number"
            id="packageSizeInput"
            placeholder="Enter size"
            value={packageSizeInput}
            onChange={(e) => setPackageSizeInput(e.target.value)}
            onKeyPress={(e) => {
              if (e.key === 'Enter') {
                addPackage()
              }
            }}
          />
          <button onClick={addPackage}>✔️</button>
        </div>
      </div>

      {/* Block 2: Order */}
      <div className="container order">
        <h3>Order</h3>
        <input
          type="number"
          id="orderSize"
          placeholder="Enter order size"
          value={orderSize}
          onChange={(e) => setOrderSize(e.target.value)}
          onKeyPress={(e) => {
            if (e.key === 'Enter') {
              calculate()
            }
          }}
        />
      </div>

      {/* Block 3: Calculate */}
      <div className="container" style={{ border: 'none' }}>
        <button className="full-width-btn" onClick={calculate}>
          Calculate
        </button>
      </div>

      {/* Block 4: Results */}
      <div className="container results">
        <h3>Results</h3>
        <table>
          <thead>
            <tr>
              <th>Package Size</th>
              <th>Count</th>
            </tr>
          </thead>
          <tbody id="results-table">
            {Object.entries(results).map(([packageSize, count]) => (
              <tr key={packageSize}>
                <td>{packageSize}</td>
                <td>{count}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>

      {/* Swagger Modal */}
      <div id="swagger-modal" className={swaggerOpen ? 'open' : ''}>
        <div className="swagger-content">
          <button className="close-btn" onClick={closeSwagger}>
            ✖
          </button>
          <iframe
            id="swagger-frame"
            src={`${API_BASE_URL}/swagger/`}
            title="Swagger UI"
          ></iframe>
        </div>
      </div>
    </>
  )
}

export default App

